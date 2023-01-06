package polygon

import (
	"common/blocktype"
	"common/client/polygonclient"
	"common/salt"
	"dao/dbmodel"
	"dao/dbsign"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"
	"sign/config"
	"sign/model"
	"sign/service"
	"strings"
	utilsCrypto "utils/utils/crypto"
	"utils/utils/kms"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"go.uber.org/zap"
)

// polygon地址生成
func newPolygonAddr() (*blocktype.Account, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	address := strings.ToLower(crypto.PubkeyToAddress(key.PublicKey).Hex())
	privateKey := hex.EncodeToString(crypto.FromECDSA(key))
	account := &blocktype.Account{
		PrivateKey: privateKey,
		Address:    address,
	}
	return account, nil
}

func CreateAccounts(params model.AccountReq) ([]string, error) {

	awsConfig := config.CONFIG.AwsConfig
	awsConfig.AccessKeyID = utilsCrypto.AesHashDecryptTool(config.CONFIG.AwsConfig.AccessKeyID, config.KMS_ACCESSKEY_PWD)
	awsConfig.SecretAccessKey = utilsCrypto.AesHashDecryptTool(config.CONFIG.AwsConfig.SecretAccessKey, config.KMS_ACCESSKEY_PWD)
	kmsClient, err := kms.GetKMSClient(awsConfig)
	if err != nil {
		config.LOG.Error("GetKMSClient err", zap.Any("err", err))
		return nil, err
	}

	dk, err := kmsClient.GenerateDataKey()
	if err != nil {
		config.LOG.Error("KMS GenerateDataKey err", zap.Any("err", err))
		return nil, err
	}

	addresses := make([]string, 0)
	count := params.Count / dbmodel.MaxLenPolygonAddr
	if count == 0 || params.Count%dbmodel.MaxLenPolygonAddr > 0 {
		count += 1
	}
	for i := 0; i < int(count); i++ {
		size := dbmodel.MaxLenPolygonAddr
		if i == int(count)-1 && params.Count%dbmodel.MaxLenPolygonAddr > 0 {
			size = int(params.Count) % dbmodel.MaxLenPolygonAddr
		}
		_addresses, err := doCreateAccounts(dk, size)
		if err != nil {
			config.LOG.Error("doCreateAccounts err", zap.Any("params.Count", params.Count), zap.Any("err", err))
			continue
		}
		addresses = append(addresses, _addresses...)
	}
	return addresses, nil
}

func doCreateAccounts(dk *kms.DataKey, count int) ([]string, error) {
	_addresses := make([]string, 0)
	accounts := make([]dbsign.PolygonAddressPool, 0)
	backup := [][]string{}

	for i := 0; i < count; i++ {
		address, err := newPolygonAddr()
		if err != nil || address.PrivateKey == "" || address.Address == "" {
			config.LOG.Error("new polygon address err", zap.Any("err", err), zap.Any("address", address))
			return _addresses, err
		}
		addressBack := []string{address.Address, address.PrivateKey}
		backup = append(backup, addressBack)

		// kms 密钥+盐
		key := string(dk.Plaintext) + salt.GetGenerateSalt()
		// 加密私钥
		priKey := utilsCrypto.AesHashEncodeTool(address.PrivateKey, key)
		if strings.EqualFold(priKey, "") {
			return nil, errors.New("aes key is null")
		}
		// 地址加签
		addrSignPriKey, err := base64.StdEncoding.DecodeString(config.CONFIG.System.AddrSignPrivateKey)
		if err != nil {
			config.LOG.Error("get addr sign private key err", zap.Any("err", err.Error()))
			return nil, err
		}
		addrSign := utilsCrypto.RsaSignWithSha256([]byte(address.Address), addrSignPriKey)
		if addrSign == "" {
			config.LOG.Error("addrSign RsaSignWithSha256 err", zap.Any("address", address.Address))
			return nil, errors.New("addr sign err")
		}
		accounts = append(accounts, dbsign.PolygonAddressPool{
			Address:    address.Address,
			PrivateKey: priKey,
			Sign:       addrSign,
			KmsKey:     base64.StdEncoding.EncodeToString(dk.CiphertextBlob),
		})
	}
	err := service.PrivateBackup(backup, "ETH")
	if err != nil {
		config.LOG.Error("private key backup error", zap.String("err", err.Error()))
		return nil, err
	}
	service.BackPwd("ETH")

	if len(accounts) != count {
		config.LOG.Error("accounts.size != count", zap.Int("accounts.size", len(accounts)),
			zap.Int("count", int(count)))
		return _addresses, errors.New("accounts.size != count")
	}
	// 更新db
	_count, err := dbsign.BatchInsertAccountsPolygon(accounts)
	if err != nil {
		config.LOG.Error("BatchInsertAccounts err", zap.Any("err", err))
		return _addresses, err
	}
	if _count != int64(count) { // 理论上不会出现
		config.LOG.Error("BatchInsertAccounts count err", zap.Any("err", err))
		return _addresses, err
	}
	for _, account := range accounts {
		_addresses = append(_addresses, account.Address)
	}
	return _addresses, nil
}

func Sign(params model.RawDataReq) (string, error) {
	rawData, _ := json.Marshal(params.Transaction)
	var rawTx types.LegacyTx
	if json.Unmarshal(rawData, &rawTx) != nil {
		config.LOG.Error("params cast err", zap.Any("params", params))
		return "", errors.New("sign params cast error")
	}
	awsConfig := config.CONFIG.AwsConfig
	awsConfig.AccessKeyID = utilsCrypto.AesHashDecryptTool(config.CONFIG.AwsConfig.AccessKeyID, config.KMS_ACCESSKEY_PWD)
	awsConfig.SecretAccessKey = utilsCrypto.AesHashDecryptTool(config.CONFIG.AwsConfig.SecretAccessKey, config.KMS_ACCESSKEY_PWD)
	// kms client
	kmsClient, err := kms.GetKMSClient(awsConfig)
	if err != nil {
		config.LOG.Error("GetKMSClient err", zap.Any("err", err))
		return "", err
	}
	// 查询地址
	account, err := dbsign.GetPolygonAccount(params.Address)
	if err != nil {
		config.LOG.Error("dbmodel.GetAccount err", zap.Any("address", params.Address), zap.Any("err", err))
		return "", err
	}
	// 用kmsSecret解密私钥
	kmsKey, _ := base64.StdEncoding.DecodeString(account.KmsKey)
	kmsSecret, err := kmsClient.Decrypt(kmsKey)
	if err != nil {
		config.LOG.Error("kmsClient.Decrypt err", zap.Any("address", params.Address), zap.Any("err", err))
		return "", err
	}
	key := string(kmsSecret) + salt.GetGenerateSalt()
	privateKey := utilsCrypto.AesHashDecryptTool(account.PrivateKey, key)
	if privateKey == "" {
		config.LOG.Error("utilsCrypto.AesHashDecryptTool is empty")
		return "", err
	}
	tx := types.NewTx(&rawTx)
	chainId, ok := config.CONFIG.System.ChainId[strings.ToLower(params.Symbol)]
	if !ok {
		config.LOG.Error("get chain id err", zap.Any("symbol", params.Symbol), zap.Any("chainids", config.CONFIG.System.ChainId))
		return "", errors.New("get chain id err " + params.Symbol)
	}
	signer := types.NewEIP155Signer(big.NewInt(chainId))
	signedTx, err := polygonclient.SignTransaction(tx, privateKey, signer)
	if err != nil {
		config.LOG.Error("SignTransaction err", zap.Any("sub_symbol", params.SubSymbol), zap.Any("err", err))
		return "", err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		config.LOG.Error("SignTransaction EncodeToBytes err", zap.Any("sub_symbol", params.SubSymbol), zap.Any("err", err))
		return "", err
	}

	config.LOG.Warn("====================", zap.Any("chain_id", chainId))
	return hexutil.Encode(data), nil
}
