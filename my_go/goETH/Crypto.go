package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//随机熵，用于加密安全
var randSign = "22220316zafes20180lk7zafes20180619zafepikas"

//随机key，用于创建公钥和私钥
var randKey = "lk0f7279c18d439459435s714797c9680335a320"

var PriKey *ecdsa.PrivateKey
var PubKey *ecdsa.PublicKey

/*
func init() {
	// 初始化生成私匙公匙
	priFile, _ := os.Create("ec-pri.pem")
	pubFile, _ := os.Create("ec-pub.pem")
	if err := generateKey(priFile, pubFile); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// 加载私匙公匙
	if err := loadKey(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
*/
func main() {

	// 预加密数据
	text := "hello dalgurak"
	// hash签名
	hashText := sha256.New().Sum([]byte(text))
	// 加载私匙公匙
	if err := loadKey(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return
	// 加密私匙
	r, s, err := ecdsa.Sign(strings.NewReader(randSign), PriKey, hashText)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// 公匙验证hash的正确性
	b := ecdsa.Verify(PubKey, hashText, r, s)
	fmt.Println(b)
}

// 生成密匙对
func generateKey(priFile, pubFile *os.File) error {
	lenth := len(randKey)
	if lenth < 224/8 {
		return errors.New("私钥长度太短，至少为36位！")
	}
	// 根据随机密匙的长度创建私匙
	var curve elliptic.Curve
	if lenth > 521/8+8 {
		curve = elliptic.P521()
	} else if lenth > 384/8+8 {
		curve = elliptic.P384()
	} else if lenth > 256/8+8 {
		curve = elliptic.P256()
	} else if lenth > 224/8+8 {
		curve = elliptic.P224()
	}
	// 生成私匙
	priKey, err := ecdsa.GenerateKey(curve, strings.NewReader(randKey))
	if err != nil {
		return err
	}
	// *****************保存私匙*******************
	// 序列化私匙
	priBytes, err := x509.MarshalECPrivateKey(priKey)
	if err != nil {
		return err
	}
	priBlock := pem.Block{
		Type:  "ECD PRIVATE KEY",
		Bytes: priBytes,
	}
	// 编码私匙,写入文件
	if err := pem.Encode(priFile, &priBlock); err != nil {
		return err
	}
	// *****************保存公匙*******************
	// 序列化公匙
	pubBytes, err := x509.MarshalPKIXPublicKey(&priKey.PublicKey)
	if err != nil {
		return err
	}
	pubBlock := pem.Block{
		Type:  "ECD PUBLIC KEY",
		Bytes: pubBytes,
	}
	// 编码公匙,写入文件
	if err := pem.Encode(pubFile, &pubBlock); err != nil {
		return err
	}
	return nil
}

// 加载私匙公匙
func loadKey() error {
	// 读取密匙
	pri, _ := ioutil.ReadFile("ec-pri.pem")
	pub, _ := ioutil.ReadFile("ec-pub.pem")
	// 解码私匙
	block, _ := pem.Decode(pri)
	var err error
	// 反序列化私匙
	PriKey, err = x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	str := hex.EncodeToString(block.Bytes)
	fmt.Printf("wch----:%+v\n %+v\n", string(block.Bytes), str)
	// 解码公匙
	block, _ = pem.Decode(pub)
	// 反序列化公匙
	var i interface{}
	i, err = x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	// PubKey = (*ecdsa.PublicKey)(i)
	var ok bool
	PubKey, ok = i.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("the public conversion error")
	}
	return nil
}
