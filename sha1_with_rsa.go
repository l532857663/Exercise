package cryptoutil

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"

	logging "github.com/op/go-logging"
)

const (
	PRIVATE_KEY_TYPE_PKCS1 = "PKCS1"
	PRIVATE_KEY_TYPE_PKCS8 = "PKCS8"
)

var (
	logger = logging.MustGetLogger("cryptoutil")
)

type PaymaxSigner interface {
	GenSignature(data []byte) (string, error)
	VerifySignature(data []byte, sign string) bool
}

type PKCSSigner struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func readKeyFile(keyFilePth string) ([]byte, error) {
	fp, err := os.Open(keyFilePth)
	if err != nil {
		return nil, err
	}
	defer fp.Close()

	key, err := ioutil.ReadAll(fp)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func DefaultPaymaxSigner(privateKeyPath, publicKeyPath, privateKeyType string) (PaymaxSigner, error) {
	return newPKCSSigner(privateKeyPath, publicKeyPath, privateKeyType)
}

func newPKCSSigner(privateKeyPath, publicKeyPath, privateKeyType string) (PaymaxSigner, error) {
	priKey, err := setPrivateKey(privateKeyPath, privateKeyType)
	if err != nil {
		logger.Errorf("SetPrivateKey error: %s", err.Error())
		return nil, err
	}

	pubKey, err := setPublicKey(publicKeyPath)
	if err != nil {
		logger.Errorf("SetPublicKey error: %s", err.Error())
		return nil, err
	}

	return &PKCSSigner{privateKey: priKey, publicKey: pubKey}, nil
}

func setPrivateKey(privateKeyPath, privateKeyType string) (*rsa.PrivateKey, error) {
	privateKey, err := readKeyFile(privateKeyPath)
	if err != nil {
		logger.Errorf("load private key file error: %s", err.Error())
		return nil, err
	}

	block, _ := pem.Decode(privateKey)
	if block == nil {
		logger.Error("pem.Decode error!!!")
		return nil, errors.New("private key is empty or format error!")
	}

	var priKey *rsa.PrivateKey

	switch privateKeyType {
	case PRIVATE_KEY_TYPE_PKCS1:
		{
			priKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				logger.Errorf("ParsePKCS1PrivateKey error: %s", err.Error())
				return nil, err
			}
		}
	case PRIVATE_KEY_TYPE_PKCS8:
		{
			prikey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				logger.Errorf("ParsePKCS8PrivateKey error: %s", err.Error())
				return nil, err
			}

			priKey = prikey.(*rsa.PrivateKey)
		}
	default:
		{
			return nil, errors.New("Unknown private key type")
		}
	}

	return priKey, nil
}

func setPublicKey(publicKeyPath string) (*rsa.PublicKey, error) {
	publicKey, err := readKeyFile(publicKeyPath)
	if err != nil {
		logger.Errorf("load public key file error: %s", err.Error())
		return nil, err
	}

	block, _ := pem.Decode(publicKey)
	if block == nil {
		logger.Error("pem.Decode error!!!")
		return nil, errors.New("publicKey key is empty or format error!")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logger.Errorf("ParsePKIXPublicKey error: %s", err.Error())
		return nil, err
	}

	return pubKey.(*rsa.PublicKey), nil
}

func (ps *PKCSSigner) GenSignature(data []byte) (string, error) {
	h := crypto.Hash.New(crypto.SHA1)
	h.Write(data)
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, ps.privateKey, crypto.SHA1, hashed)
	if err != nil {
		logger.Errorf("RSA sign SignPKCS1v15 error: %s", err.Error())
		return "", err
	}

	b64Signature := base64.StdEncoding.EncodeToString(signature)

	return b64Signature, nil
}

func (ps *PKCSSigner) VerifySignature(data []byte, sign string) bool {
	h := crypto.Hash.New(crypto.SHA1)
	h.Write(data)
	hashed := h.Sum(nil)

	parseSign, _ := base64.StdEncoding.DecodeString(sign)

	err := rsa.VerifyPKCS1v15(ps.publicKey, crypto.SHA1, hashed, parseSign)
	if err != nil {
		logger.Errorf("RSA verify VerifyPKCS1v15 error: %s", err.Error())
		return false
	}

	return true
}
