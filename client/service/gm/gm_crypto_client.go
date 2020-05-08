package gm

import (
	"crypto/ecdsa"
	"crypto/rand"

	"github.com/xuperchain/crypto/gm/account"
	"github.com/xuperchain/crypto/gm/config"
	"github.com/xuperchain/crypto/gm/gmsm/sm2"
	"github.com/xuperchain/crypto/gm/hash"
	"github.com/xuperchain/crypto/gm/hdwallet/key"
	"github.com/xuperchain/crypto/gm/sign"

	walletRand "github.com/xuperchain/crypto/core/hdwallet/rand"
)

type GmCryptoClient struct {
}

// --- 哈希算法相关 start ---

// 使用SHA256做单次哈希运算
func (gcc *GmCryptoClient) HashUsingSM3(data []byte) []byte {
	hashResult := hash.HashUsingSM3(data)
	return hashResult
}

// 使用Hmac512做单次哈希运算
func (gcc *GmCryptoClient) HashUsingHmac512(data, key []byte) []byte {
	hashResult := hash.HashUsingHmac512(data, key)
	return hashResult
}

// 使用Ripemd160做单次哈希运算
func (gcc *GmCryptoClient) HashUsingRipemd160(data []byte) []byte {
	hashResult := hash.HashUsingRipemd160(data)
	return hashResult
}

// --- 哈希算法相关 end ---

// --- 随机数相关 start ---

// 产生随机熵
func (gcc *GmCryptoClient) GenerateEntropy(bitSize int) ([]byte, error) {
	entropyByte, err := walletRand.GenerateEntropy(bitSize)
	return entropyByte, err
}

// --- 随机数相关 end ---

// --- 助记词相关 start ---

// 将随机熵转为助记词
func (gcc *GmCryptoClient) GenerateMnemonic(entropy []byte, language int) (string, error) {
	mnemonic, err := walletRand.GenerateMnemonic(entropy, language)
	return mnemonic, err
}

// 将助记词转为指定长度的随机数种子，在此过程中，校验助记词是否合法
func (gcc *GmCryptoClient) GenerateSeedWithErrorChecking(mnemonic string, password string, keyLen int, language int) ([]byte, error) {
	seed, err := walletRand.GenerateSeedWithErrorChecking(mnemonic, password, keyLen, language)
	return seed, err
}

// --- 助记词相关 end ---

// --- 密钥字符串转换相关 start ---

// 获取ECC私钥的json格式的表达
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPrivateKeyJsonFormat, err := account.GetEcdsaPrivateKeyJsonFormat(k)
	return jsonEcdsaPrivateKeyJsonFormat, err
}

// 获取ECC公钥的json格式的表达
func (gcc *GmCryptoClient) GetEcdsaPublicKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPublicKeyJsonFormat, err := account.GetEcdsaPublicKeyJsonFormat(k)
	return jsonEcdsaPublicKeyJsonFormat, err
}

// 从json格式私钥内容字符串产生ECC私钥
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyFromJsonStr(keyStr string) (*ecdsa.PrivateKey, error) {
	jsonBytes := []byte(keyStr)
	return account.GetEcdsaPrivateKeyFromJson(jsonBytes)
}

// 从json格式公钥内容字符串产生ECC公钥
func (gcc *GmCryptoClient) GetEcdsaPublicKeyFromJsonStr(keyStr string) (*ecdsa.PublicKey, error) {
	jsonBytes := []byte(keyStr)
	return account.GetEcdsaPublicKeyFromJson(jsonBytes)
}

// --- 密钥字符串转换相关 end ---

// --- 地址生成相关 start ---

// 使用单个公钥来生成钱包地址
func (gcc *GmCryptoClient) GetAddressFromPublicKey(key *ecdsa.PublicKey) (string, error) {
	address, err := account.GetAddressFromPublicKey(key)
	return address, err
}

// 使用多个公钥来生成钱包地址（环签名，多重签名地址）
func (gcc *GmCryptoClient) GetAddressFromPublicKeys(keys []*ecdsa.PublicKey) (string, error) {
	address, err := account.GetAddressFromPublicKeys(keys)
	return address, err
}

// 验证钱包地址是否是合法的格式。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (gcc *GmCryptoClient) CheckAddressFormat(address string) (bool, uint8) {
	isValid, nVersion := account.CheckAddressFormat(address)
	return isValid, nVersion
}

// 验证钱包地址是否和指定的公钥match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (gcc *GmCryptoClient) VerifyAddressUsingPublicKey(address string, pub *ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := account.VerifyAddressUsingPublicKey(address, pub)
	return isValid, nVersion
}

// 验证钱包地址（环签名，多重签名地址）是否和指定的公钥数组match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (gcc *GmCryptoClient) VerifyAddressUsingPublicKeys(address string, pub []*ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := account.VerifyAddressUsingPublicKeys(address, pub)
	return isValid, nVersion
}

// --- 地址生成相关 end ---

// --- 账户相关 start ---

func (gcc *GmCryptoClient) GenerateKeyBySeed(seed []byte) (*ecdsa.PrivateKey, error) {
	curve := sm2.P256Sm2()
	privateKey, err := sign.GenerateKeyBySeed(curve, seed)
	return privateKey, err
}

// ExportNewAccount 创建新账户(不使用助记词，不推荐使用)
func (gcc *GmCryptoClient) ExportNewAccount(path string) error {
	privateKey, err := ecdsa.GenerateKey(sm2.P256Sm2(), rand.Reader)
	if err != nil {
		return err
	}
	return account.ExportNewAccount(path, privateKey)
}

// 创建含有助记词的新的账户，返回的字段：（助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (gcc *GmCryptoClient) CreateNewAccountWithMnemonic(language int, strength uint8) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := account.CreateNewAccountWithMnemonic(language, strength, config.Gm)
	return ecdsaAccount, err
}

// 创建新的账户，并用支付密码加密私钥后存在本地，
// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (gcc *GmCryptoClient) CreateNewAccountAndSaveSecretKey(path string, language int, strength uint8, password string) (*account.ECDSAInfo, error) {
	ecdasaInfo, err := key.CreateAndSaveSecretKey(path, walletRand.SimplifiedChinese, account.StrengthHard, password, config.Gm)
	return ecdasaInfo, err
}

// 创建新的账户，并导出相关文件（含助记词）到本地。生成如下几个文件：1.助记词，2.私钥，3.公钥，4.钱包地址
func (gcc *GmCryptoClient) ExportNewAccountWithMnemonic(path string, language int, strength uint8) error {
	err := account.ExportNewAccountWithMnemonic(path, language, strength, config.Gm)
	return err
}

// 从助记词恢复钱包账户
// TODO: 后续可以从助记词中识别出语言类型
func (gcc *GmCryptoClient) RetrieveAccountByMnemonic(mnemonic string, language int) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := account.GenerateAccountByMnemonic(mnemonic, language)
	return ecdsaAccount, err
}

// 从助记词恢复钱包账户，并用支付密码加密私钥后存在本地，
// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (gcc *GmCryptoClient) RetrieveAccountByMnemonicAndSavePrivKey(path string, language int, mnemonic string, password string) (*account.ECDSAInfo, error) {
	ecdsaAccount, err := key.CreateAndSaveSecretKeyWithMnemonic(path, language, mnemonic, password)
	return ecdsaAccount, err
}

// 使用支付密码加密账户信息并返回加密后的数据（后续用来回传至云端）
func (gcc *GmCryptoClient) EncryptAccount(info *account.ECDSAAccount, password string) (*account.ECDSAAccountToCloud, error) {
	ecdsaAccountToCloud, err := key.EncryptAccount(info, password)
	return ecdsaAccountToCloud, err
}

// 从导出的私钥文件读取私钥的byte格式
func (gcc *GmCryptoClient) GetBinaryEcdsaPrivateKeyFromFile(path string, password string) ([]byte, error) {
	binaryEcdsaPrivateKey, err := key.GetBinaryEcdsaPrivateKeyFromFile(path, password)
	return binaryEcdsaPrivateKey, err
}

// 使用支付密码从导出的私钥文件读取私钥
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyFromFileByPassword(path string, password string) (*ecdsa.PrivateKey, error) {
	ecdsaPrivateKey, err := key.GetEcdsaPrivateKeyFromFile(path, password)
	return ecdsaPrivateKey, err
}

// 使用支付密码从二进制加密字符串获取真实私钥的字节数组
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyBytesFromEncryptedStringByPassword(encryptedPrivateKey string, password string) ([]byte, error) {
	binaryEcdsaPrivateKey, err := key.GetBinaryEcdsaPrivateKeyFromString(encryptedPrivateKey, password)
	return binaryEcdsaPrivateKey, err
}

// 使用支付密码从二进制加密字符串获取真实ECC私钥
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyFromEncryptedStringByPassword(encryptedPrivateKey string, password string) (*ecdsa.PrivateKey, error) {
	binaryEcdsaPrivateKey, err := key.GetEcdsaPrivateKeyFromString(encryptedPrivateKey, password)
	return binaryEcdsaPrivateKey, err
}

// 从导出的私钥文件读取私钥
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	ecdsaPrivateKey, err := account.GetEcdsaPrivateKeyFromFile(filename)
	return ecdsaPrivateKey, err
}

// 从导出的公钥文件读取公钥
func (gcc *GmCryptoClient) GetEcdsaPublicKeyFromFile(filename string) (*ecdsa.PublicKey, error) {
	ecdsaPublicKey, err := account.GetEcdsaPublicKeyFromFile(filename)
	return ecdsaPublicKey, err
}

// --- 账户相关 end ---

// --- 普通单签名相关 start ---

// 使用ECC私钥来签名
func (gcc *GmCryptoClient) SignECDSA(k *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	signature, err := sign.SignECDSA(k, msg)
	return signature, err
}

// 使用ECC私钥来签名，生成统一签名的新签名函数
func (gcc *GmCryptoClient) SignV2ECDSA(k *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	signature, err := sign.SignV2ECDSA(k, msg)
	return signature, err
}

// 使用ECC公钥来验证签名，验证统一签名的新签名函数
func (gcc *GmCryptoClient) VerifyECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyECDSA(k, signature, msg)
	return result, err
}

// 使用ECC公钥来验证签名，验证统一签名的新签名函数 -- 供统一验签函数调用
func (gcc *GmCryptoClient) VerifyV2ECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyV2ECDSA(k, signature, msg)
	return result, err
}

// --- 普通单签名相关 end ---
