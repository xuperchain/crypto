package gm

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/xuperchain/crypto/gm/account"
	"github.com/xuperchain/crypto/gm/config"
	"github.com/xuperchain/crypto/gm/gmsm/sm2"
	"github.com/xuperchain/crypto/gm/hash"
	"github.com/xuperchain/crypto/gm/sign"

	walletRand "github.com/xuperchain/crypto/core/hdwallet/rand"
)

type GmCryptoClient struct {
}

// --- 哈希算法相关 start ---

// 使用SHA256做单次哈希运算
func (xcc *GmCryptoClient) HashUsingSM3(data []byte) []byte {
	hashResult := hash.HashUsingSM3(data)
	return hashResult
}

// 使用Hmac512做单次哈希运算
func (xcc *GmCryptoClient) HashUsingHmac512(data, key []byte) []byte {
	hashResult := hash.HashUsingHmac512(data, key)
	return hashResult
}

// 使用Ripemd160做单次哈希运算
func (xcc *GmCryptoClient) HashUsingRipemd160(data []byte) []byte {
	hashResult := hash.HashUsingRipemd160(data)
	return hashResult
}

// --- 哈希算法相关 end ---

// --- 随机数相关 start ---

// 产生随机熵
func (xcc *GmCryptoClient) GenerateEntropy(bitSize int) ([]byte, error) {
	entropyByte, err := walletRand.GenerateEntropy(bitSize)
	return entropyByte, err
}

// --- 随机数相关 end ---

// --- 助记词相关 start ---

// 将随机熵转为助记词
func (xcc *GmCryptoClient) GenerateMnemonic(entropy []byte, language int) (string, error) {
	mnemonic, err := walletRand.GenerateMnemonic(entropy, language)
	return mnemonic, err
}

// 将助记词转为指定长度的随机数种子，在此过程中，校验助记词是否合法
func (xcc *GmCryptoClient) GenerateSeedWithErrorChecking(mnemonic string, password string, keyLen int, language int) ([]byte, error) {
	seed, err := walletRand.GenerateSeedWithErrorChecking(mnemonic, password, keyLen, language)
	return seed, err
}

// --- 助记词相关 end ---

// --- 密钥字符串转换相关 start ---

// 获取ECC私钥的json格式的表达
func (xcc *GmCryptoClient) GetEcdsaPrivateKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPrivateKeyJsonFormat, err := account.GetEcdsaPrivateKeyJsonFormat(k)
	return jsonEcdsaPrivateKeyJsonFormat, err
}

// 获取ECC公钥的json格式的表达
func (xcc *GmCryptoClient) GetEcdsaPublicKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPublicKeyJsonFormat, err := account.GetEcdsaPublicKeyJsonFormat(k)
	return jsonEcdsaPublicKeyJsonFormat, err
}

// --- 密钥字符串转换相关 end ---

// --- 地址生成相关 start ---

// 使用单个公钥来生成钱包地址
func (xcc *GmCryptoClient) GetAddressFromPublicKey(key *ecdsa.PublicKey) (string, error) {
	address, err := account.GetAddressFromPublicKey(key)
	return address, err
}

// 使用多个公钥来生成钱包地址（环签名，多重签名地址）
func (xcc *GmCryptoClient) GetAddressFromPublicKeys(keys []*ecdsa.PublicKey) (string, error) {
	address, err := account.GetAddressFromPublicKeys(keys)
	return address, err
}

// 验证钱包地址是否是合法的格式。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *GmCryptoClient) CheckAddressFormat(address string) (bool, uint8) {
	isValid, nVersion := account.CheckAddressFormat(address)
	return isValid, nVersion
}

// 验证钱包地址是否和指定的公钥match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *GmCryptoClient) VerifyAddressUsingPublicKey(address string, pub *ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := account.VerifyAddressUsingPublicKey(address, pub)
	return isValid, nVersion
}

// 验证钱包地址（环签名，多重签名地址）是否和指定的公钥数组match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *GmCryptoClient) VerifyAddressUsingPublicKeys(address string, pub []*ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := account.VerifyAddressUsingPublicKeys(address, pub)
	return isValid, nVersion
}

// --- 地址生成相关 end ---

// --- 账户相关 start ---

func (xcc *GmCryptoClient) GenerateKeyBySeed(seed []byte) (*ecdsa.PrivateKey, error) {
	curve := sm2.P256Sm2()
	privateKey, err := sign.GenerateKeyBySeed(curve, seed)
	return privateKey, err
}

// ExportNewAccount 创建新账户(不使用助记词，不推荐使用)
func (xcc *GmCryptoClient) ExportNewAccount(path string) error {
	privateKey, err := ecdsa.GenerateKey(sm2.P256Sm2(), rand.Reader)
	if err != nil {
		return err
	}
	return account.ExportNewAccount(path, privateKey)
}

// 创建含有助记词的新的账户，返回的字段：（助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (xcc *GmCryptoClient) CreateNewAccountWithMnemonic(language int, strength uint8) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := account.CreateNewAccountWithMnemonic(language, strength, config.Nist)
	return ecdsaAccount, err
}

// --- 账户相关 end ---
