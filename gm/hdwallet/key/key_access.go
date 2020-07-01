package key

import (
	"crypto/ecdsa"
	//	"crypto/elliptic"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xuperchain/crypto/common/account"
	"github.com/xuperchain/crypto/gm/ecies"
	"github.com/xuperchain/crypto/gm/gmsm/sm2"
	"github.com/xuperchain/crypto/gm/hash"

	accountUtil "github.com/xuperchain/crypto/gm/account"
	aesUtil "github.com/xuperchain/crypto/gm/aes"
)

func GetBinaryEcdsaPrivateKeyFromFile(path string, password string) ([]byte, error) {
	filename := path + "private.key"
	content, err := readFileUsingFilename(filename)
	if err != nil {
		log.Printf("readFileUsingFilename failed, the err is %v", err)
		return nil, err
	}

	// 将aes对称加密的密钥扩展至32字节
	//	newPassword := hash.DoubleSha256([]byte(password))
	newPassword := hash.HashUsingSM3([]byte(password))

	originalContent, err := aesUtil.Decrypt(content, newPassword)
	if err != nil {
		log.Printf("Decrypt private key file failed, the err is %v", err)
		return nil, err
	}

	return originalContent, nil
}

// GetBinaryEcdsaPrivateKeyFromString通过二进制字符串获取真实私钥的字节数组
func GetBinaryEcdsaPrivateKeyFromString(encryptPrivateKey string, password string) ([]byte, error) {
	// 将aes对称加密的密钥扩展至32字节
	//	newPassword := hash.DoubleSha256([]byte(password))
	newPassword := hash.HashUsingSM3([]byte(password))

	originalContent, err := aesUtil.Decrypt([]byte(encryptPrivateKey), newPassword)
	if err != nil {
		log.Printf("Decrypt private key string failed, the err is %v", err)
		return nil, err
	}

	return originalContent, nil
}

// GetEcdsaPrivateKeyFromString通过二进制字符串获取真实私钥
func GetEcdsaPrivateKeyFromString(encryptPrivateKey string, password string) (*ecdsa.PrivateKey, error) {
	originalContent, err := GetBinaryEcdsaPrivateKeyFromString(encryptPrivateKey, password)
	if err != nil {
		log.Printf("Decrypt private key string failed, the err is %v", err)
		return nil, err
	}

	return accountUtil.GetEcdsaPrivateKeyFromJson(originalContent)
}

func GetEcdsaPrivateKeyFromFile(path string, password string) (*ecdsa.PrivateKey, error) {
	originalContent, err := GetBinaryEcdsaPrivateKeyFromFile(path, password)
	if err != nil {
		log.Printf("GetBinaryEcdsaPrivateKeyFromFile failed, the err is %v", err)
		return nil, err
	}

	return accountUtil.GetEcdsaPrivateKeyFromJson(originalContent)
}

func readFileUsingFilename(filename string) ([]byte, error) {
	contentBytes, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		log.Printf("File[%v] does not exist", filename)
	}
	if err != nil {
		return nil, err
	}
	content, err := base64.StdEncoding.DecodeString(string(contentBytes))
	if err != nil {
		return nil, err
	}

	return content, err
}

func GetEncryptedPrivateKeyFromFile(path string) (string, error) {
	filename := path + "private.key"
	content, err := readFileUsingFilename(filename)
	if err != nil {
		log.Printf("readFileUsingFilename failed, the err is %v", err)
		return "", err
	}
	return string(content), nil
}

func GetEcdsaPublicKeyFromJson(jsonContent []byte) (*ecdsa.PublicKey, error) {
	publicKey := new(accountUtil.ECDSAPublicKey)
	err := json.Unmarshal(jsonContent, publicKey)
	if err != nil {
		return nil, err //json有问题
	}
	if publicKey.Curvname != "SM2-P-256" {
		log.Printf("curve [%v] is not supported yet.", publicKey.Curvname)
		err = fmt.Errorf("curve [%v] is not supported yet.", publicKey.Curvname)
		return nil, err
	}
	ecdsaPublicKey := &ecdsa.PublicKey{}
	//	ecdsaPublicKey.Curve = elliptic.P256()
	ecdsaPublicKey.Curve = sm2.P256Sm2()
	ecdsaPublicKey.X = publicKey.X
	ecdsaPublicKey.Y = publicKey.Y
	return ecdsaPublicKey, nil
}

// GetAccountFromLocal 读取本地文件获取账户信息
func GetAccountFromLocal(path string) (*account.ECDSAAccountToCloud, error) {
	account := new(account.ECDSAAccountToCloud)
	privateKeyFile := path + "private.key"
	privateKey, err := readFileUsingFilename(privateKeyFile)
	if err != nil {
		log.Printf("readFileUsingFilename failed, the err is %v", err)
		return nil, err
	}

	addressFile := path + "address"
	address, err := readFileUsingFilename(addressFile)
	if err != nil {
		log.Printf("readFileUsingFilename failed, the err is %v", err)
		return nil, err
	}
	account.JsonEncryptedPrivateKey = string(privateKey)
	account.Address = string(address)

	return account, nil
}

// EncryptByKey 加密
func EncryptByKey(info string, key string) (string, error) {
	// 将aes对称加密的密钥扩展至32字节
	//	newPassword := hash.DoubleSha256([]byte(key))
	newPassword := hash.HashUsingSM3([]byte(key))

	// 加密info
	cipherInfo, err := aesUtil.Encrypt([]byte(info), newPassword)
	if err != nil {
		return "", err
	}
	return string(cipherInfo), err
}

// DecryptByKey 解密
func DecryptByKey(cipherInfo string, key string) (string, error) {
	// 将aes对称加密的密钥扩展至32字节
	//	newPassword := hash.DoubleSha256([]byte(key))
	newPassword := hash.HashUsingSM3([]byte(key))

	// 解密cipherInfo
	info, err := aesUtil.Decrypt([]byte(cipherInfo), newPassword)
	if err != nil {
		return "", err
	}
	return string(info), nil
}

// GetPublicKeyByPrivateKey通过私钥获取公钥
func GetPublicKeyByPrivateKey(binaryPrivateKey string) (string, error) {
	privatekey, err := accountUtil.GetEcdsaPrivateKeyFromJson([]byte(binaryPrivateKey))
	if err != nil {
		return "", err
	}

	// 补充公钥
	jsonPublicKey, err := accountUtil.GetEcdsaPublicKeyJsonFormat(privatekey)
	if err != nil {
		return "", err
	}
	return jsonPublicKey, nil
}

// EciesEncryptByJsonPublicKey 使用字符串公钥进行ecies加密
func EciesEncryptByJsonPublicKey(publicKey string, msg string) (string, error) {
	apiPublicKey, err := GetEcdsaPublicKeyFromJson([]byte(publicKey))
	if err != nil {
		return "", errors.New("api public key is wrong")
	}
	cipherInfo, err := ecies.Encrypt(apiPublicKey, []byte(msg))
	if err != nil {
		return "", ErrParam
	}
	return string(cipherInfo), nil
}

// EciesDecryptByJsonPublicKey 使用字符串私钥进行ecies解密
func EciesDecryptByJsonPrivateKey(privateKey string, cipherInfo string) (string, error) {
	apiPrivateKey, err := accountUtil.GetEcdsaPrivateKeyFromJson([]byte(privateKey))
	if err != nil {
		return "", errors.New("api public key is wrong")
	}
	msg, err := ecies.Decrypt(apiPrivateKey, []byte(cipherInfo))
	if err != nil {
		return "", ErrParam
	}
	return string(msg), nil
}
