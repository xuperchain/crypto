package gm

import (
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"

	"github.com/xuperchain/crypto/client/service/base"
	"github.com/xuperchain/crypto/common/account"
	"github.com/xuperchain/crypto/gm/config"
	"github.com/xuperchain/crypto/gm/ecies"
	"github.com/xuperchain/crypto/gm/gmsm/sm2"
	"github.com/xuperchain/crypto/gm/hash"
	"github.com/xuperchain/crypto/gm/hdwallet/key"
	"github.com/xuperchain/crypto/gm/multisign"
	"github.com/xuperchain/crypto/gm/schnorr_ring_sign"
	"github.com/xuperchain/crypto/gm/schnorr_sign"
	"github.com/xuperchain/crypto/gm/secret_share/complex_secret_share"
	"github.com/xuperchain/crypto/gm/sign"
	"github.com/xuperchain/crypto/gm/signature"

	"github.com/xuperchain/crypto/common/utils"

	accountUtil "github.com/xuperchain/crypto/gm/account"
	aesUtil "github.com/xuperchain/crypto/gm/aes"
	hd "github.com/xuperchain/crypto/gm/hdwallet/api"
	walletRand "github.com/xuperchain/crypto/gm/hdwallet/rand"
)

type GmCryptoClient struct {
	base.CryptoClient
}

// --- 哈希算法相关 start ---

// 使用SHA256做单次哈希运算
func (gcc *GmCryptoClient) HashUsingSM3(data []byte) []byte {
	hashResult := hash.HashUsingSM3(data)
	return hashResult
}

// 使用Hmac512做哈希运算
func (gcc *GmCryptoClient) HashUsingHmac512(data, key []byte) []byte {
	hashResult := hash.HashUsingHmac512(data, key)
	return hashResult
}

// 使用Ripemd160做哈希运算
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
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyJsonFormatStr(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPrivateKeyJsonFormat, err := accountUtil.GetEcdsaPrivateKeyJsonFormat(k)
	return jsonEcdsaPrivateKeyJsonFormat, err
}

// 通过私钥获取ECC公钥的json格式的表达
func (gcc *GmCryptoClient) GetEcdsaPublicKeyJsonFormatStr(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPublicKeyJsonFormat, err := accountUtil.GetEcdsaPublicKeyJsonFormat(k)
	return jsonEcdsaPublicKeyJsonFormat, err
}

// 通过公钥获取ECC公钥的json格式的表达的字符串
func (gcc *GmCryptoClient) GetEcdsaPublicKeyJsonFormatStrFromPublicKey(k *ecdsa.PublicKey) (string, error) {
	jsonEcdsaPublicKeyJsonFormat, err := accountUtil.GetEcdsaPublicKeyJsonFormatFromPublicKey(k)
	return jsonEcdsaPublicKeyJsonFormat, err
}

// 从json格式私钥内容字符串产生ECC私钥
func (gcc *GmCryptoClient) GetEcdsaPrivateKeyFromJsonStr(keyStr string) (*ecdsa.PrivateKey, error) {
	jsonBytes := []byte(keyStr)
	return accountUtil.GetEcdsaPrivateKeyFromJson(jsonBytes)
}

// 从json格式公钥内容字符串产生ECC公钥
func (gcc *GmCryptoClient) GetEcdsaPublicKeyFromJsonStr(keyStr string) (*ecdsa.PublicKey, error) {
	jsonBytes := []byte(keyStr)
	return accountUtil.GetEcdsaPublicKeyFromJson(jsonBytes)
}

// --- 密钥字符串转换相关 end ---

// --- 地址生成相关 start ---

// 使用单个公钥来生成钱包地址
func (gcc *GmCryptoClient) GetAddressFromPublicKey(key *ecdsa.PublicKey) (string, error) {
	address, err := accountUtil.GetAddressFromPublicKey(key)
	return address, err
}

// 使用多个公钥来生成钱包地址（环签名，多重签名地址）
func (gcc *GmCryptoClient) GetAddressFromPublicKeys(keys []*ecdsa.PublicKey) (string, error) {
	address, err := accountUtil.GetAddressFromPublicKeys(keys)
	return address, err
}

// 验证钱包地址是否是合法的格式。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (gcc *GmCryptoClient) CheckAddressFormat(address string) (bool, uint8) {
	isValid, nVersion := accountUtil.CheckAddressFormat(address)
	return isValid, nVersion
}

// 验证钱包地址是否和指定的公钥match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (gcc *GmCryptoClient) VerifyAddressUsingPublicKey(address string, pub *ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := accountUtil.VerifyAddressUsingPublicKey(address, pub)
	return isValid, nVersion
}

// 验证钱包地址（环签名，多重签名地址）是否和指定的公钥数组match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (gcc *GmCryptoClient) VerifyAddressUsingPublicKeys(address string, pub []*ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := accountUtil.VerifyAddressUsingPublicKeys(address, pub)
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
	return accountUtil.ExportNewAccount(path, privateKey)
}

// 创建含有助记词的新的账户，返回的字段：（助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (gcc *GmCryptoClient) CreateNewAccountWithMnemonic(language int, strength uint8) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := accountUtil.CreateNewAccountWithMnemonic(language, strength, config.Gm)
	return ecdsaAccount, err
}

// 创建新的账户，并用支付密码加密私钥后存在本地，
// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (gcc *GmCryptoClient) CreateNewAccountAndSaveSecretKey(path string, language int, strength uint8, password string) (*account.ECDSAInfo, error) {
	ecdasaInfo, err := key.CreateAndSaveSecretKey(path, walletRand.SimplifiedChinese, accountUtil.StrengthHard, password, config.Gm)
	return ecdasaInfo, err
}

// 创建新的账户，并导出相关文件（含助记词）到本地。生成如下几个文件：1.助记词，2.私钥，3.公钥，4.钱包地址
func (gcc *GmCryptoClient) ExportNewAccountWithMnemonic(path string, language int, strength uint8) error {
	err := accountUtil.ExportNewAccountWithMnemonic(path, language, strength, config.Gm)
	return err
}

// 从助记词恢复钱包账户
// TODO: 后续可以从助记词中识别出语言类型
func (gcc *GmCryptoClient) RetrieveAccountByMnemonic(mnemonic string, language int) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := accountUtil.GenerateAccountByMnemonic(mnemonic, language)
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
	ecdsaPrivateKey, err := accountUtil.GetEcdsaPrivateKeyFromFile(filename)
	return ecdsaPrivateKey, err
}

// 从导出的公钥文件读取公钥
func (gcc *GmCryptoClient) GetEcdsaPublicKeyFromFile(filename string) (*ecdsa.PublicKey, error) {
	ecdsaPublicKey, err := accountUtil.GetEcdsaPublicKeyFromFile(filename)
	return ecdsaPublicKey, err
}

// 切分账户私钥
func (gcc *GmCryptoClient) SplitPrivateKey(jsonPrivateKey string, totalShareNumber, minimumShareNumber int) ([]string, error) {
	jsonPrivateKeyShares, err := accountUtil.SplitPrivateKey(jsonPrivateKey, totalShareNumber, minimumShareNumber)
	return jsonPrivateKeyShares, err
}

// 通过私钥片段恢复私钥
func (gcc *GmCryptoClient) RetrievePrivateKeyByShares(jsonPrivateKeyShares []string) (string, error) {
	jsonPrivateKey, err := accountUtil.RetrievePrivateKeyByShares(jsonPrivateKeyShares)
	return jsonPrivateKey, err
}

// 将私钥的曲线转化为secp256k1，并重新计算包含的公钥
func (gcc *GmCryptoClient) ChangePrivCurveToS256k1(key *ecdsa.PrivateKey) *ecdsa.PrivateKey {
	newPriv := utils.ChangePrivCurveToS256k1(key)
	return newPriv
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

// 使用ECC公钥来验证签名 -- 对应SignECDSA
func (gcc *GmCryptoClient) VerifyECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyECDSA(k, signature, msg)
	return result, err
}

// 使用ECC公钥来验证签名，验证统一签名的新签名函数  -- 内部函数，供统一验签函数调用
func (gcc *GmCryptoClient) VerifyV2ECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyV2ECDSA(k, signature, msg)
	return result, err
}

// --- 普通单签名相关 end ---

// --- 加解密相关 start ---

// 使用椭圆曲线非对称加密
func (gcc *GmCryptoClient) EncryptByEcdsaKey(publicKey *ecdsa.PublicKey, msg []byte) (cypherText []byte, err error) {
	cypherText, err = ecies.Encrypt(publicKey, msg)
	return cypherText, err
}

// 使用椭圆曲线非对称解密
func (gcc *GmCryptoClient) DecryptByEcdsaKey(privateKey *ecdsa.PrivateKey, cypherText []byte) (msg []byte, err error) {
	msg, err = ecies.Decrypt(privateKey, cypherText)
	return msg, err
}

// 使用AES对称加密算法加密
func (gcc *GmCryptoClient) EncryptByAESKey(info string, cypherKey string) (string, error) {
	cipherInfo, err := aesUtil.Encrypt([]byte(info), []byte(cypherKey))
	if err != nil {
		return "", err
	}

	return string(cipherInfo), err
}

// 使用AES对称加密算法解密
func (gcc *GmCryptoClient) DecryptByAESKey(cipherInfo string, cypherKey string) (string, error) {
	info, err := aesUtil.Decrypt([]byte(cipherInfo), []byte(cypherKey))
	if err != nil {
		return "", err
	}

	return string(info), nil
}

// 使用AES对称加密算法加密，密钥会被增强拓展，提升破解难度
func (gcc *GmCryptoClient) EncryptHardenByAESKey(info string, cypherKey string) (string, error) {
	return key.EncryptByKey(info, cypherKey)
}

// 使用AES对称加密算法解密，密钥曾经被增强拓展，提升破解难度
func (gcc *GmCryptoClient) DecryptHardenByAESKey(cipherInfo string, cypherKey string) (string, error) {
	return key.DecryptByKey(cipherInfo, cypherKey)
}

// 将经过支付密码加密的账户保存到文件中
func (gcc *GmCryptoClient) SaveEncryptedAccountToFile(account *account.ECDSAAccountToCloud, path string) error {
	return key.SaveAccountFile(account, path)
}

// --- 加解密相关 end ---

// --- 多重签名相关 start ---

// 每个多重签名算法流程的参与节点生成32位长度的随机byte，返回值可以认为是k
func (gcc *GmCryptoClient) GetRandom32Bytes() ([]byte, error) {
	return multisign.GetRandom32Bytes()
}

// 每个多重签名算法流程的参与节点生成Ri = Ki*G
func (gcc *GmCryptoClient) GetRiUsingRandomBytes(key *ecdsa.PublicKey, k []byte) []byte {
	return multisign.GetRiUsingRandomBytes(key, k)
}

// 负责计算多重签名的节点来收集所有节点的Ri，并计算R = k1*G + k2*G + ... + kn*G
func (gcc *GmCryptoClient) GetRUsingAllRi(key *ecdsa.PublicKey, arrayOfRi [][]byte) []byte {
	return multisign.GetRUsingAllRi(key, arrayOfRi)
}

// 负责计算多重签名的节点来收集所有节点的公钥Pi，并计算公共公钥：C = P1 + P2 + ... + Pn
func (gcc *GmCryptoClient) GetSharedPublicKeyForPublicKeys(keys []*ecdsa.PublicKey) ([]byte, error) {
	return multisign.GetSharedPublicKeyForPublicKeys(keys)
}

// 负责计算多重签名的节点将计算出的R和C分别传递给各个参与节点后，由各个参与节点再次计算自己的Si
// 计算 Si = Ki + HASH(C,R,m) * Xi
// X代表大数D，也就是私钥的关键参数
func (gcc *GmCryptoClient) GetSiUsingKCRM(key *ecdsa.PrivateKey, k []byte, c []byte, r []byte, message []byte) []byte {
	return multisign.GetSiUsingKCRM(key, k, c, r, message)
}

// 负责计算多重签名的节点来收集所有节点的Si，并计算出S = sum(si)
func (gcc *GmCryptoClient) GetSUsingAllSi(arrayOfSi [][]byte) []byte {
	return multisign.GetSUsingAllSi(arrayOfSi)
}

// 负责计算多重签名的节点，最终生成多重签名的统一签名格式XuperSignature
func (gcc *GmCryptoClient) GenerateMultiSignSignature(s []byte, r []byte) ([]byte, error) {
	return multisign.GenerateMultiSignSignature(s, r)
}

// 使用ECC公钥数组来进行多重签名的验证  -- 内部函数，供统一验签函数调用
func (gcc *GmCryptoClient) VerifyMultiSig(keys []*ecdsa.PublicKey, signature, message []byte) (bool, error) {
	return multisign.VerifyMultiSig(keys, signature, message)
}

// -- 多重签名的另一种用法，适用于完全中心化的流程
// 使用ECC私钥数组来进行多重签名，生成统一签名格式XuperSignature
func (gcc *GmCryptoClient) MultiSign(keys []*ecdsa.PrivateKey, message []byte) ([]byte, error) {
	return multisign.MultiSign(keys, message)
}

// --- 多重签名相关 end ---

// --- 	schnorr签名算法相关 start ---

// schnorr签名算法 生成统一签名XuperSignature
func (gcc *GmCryptoClient) SignSchnorr(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	return schnorr_sign.Sign(privateKey, message)
}

// schnorr签名算法 验证签名  -- 内部函数，供统一验签函数调用
func (gcc *GmCryptoClient) VerifySchnorr(publicKey *ecdsa.PublicKey, sig, message []byte) (bool, error) {
	return schnorr_sign.Verify(publicKey, sig, message)
}

// --- 	schnorr签名算法相关 end ---

// --- 	schnorr 环签名算法相关 start ---

// schnorr环签名算法 生成统一签名XuperSignature
func (gcc *GmCryptoClient) SignSchnorrRing(keys []*ecdsa.PublicKey, privateKey *ecdsa.PrivateKey, message []byte) (ringSignature []byte, err error) {
	return schnorr_ring_sign.Sign(keys, privateKey, message)
}

// schnorr环签名算法 验证签名  -- 内部函数，供统一验签函数调用
func (gcc *GmCryptoClient) VerifySchnorrRing(keys []*ecdsa.PublicKey, sig, message []byte) (bool, error) {
	return schnorr_ring_sign.Verify(keys, sig, message)
}

// --- 	schnorr 环签名算法相关 end ---

// --- XuperSignature 统一签名相关 start ---

// --- 统一验签算法，可以对用各种签名算法生成的签名进行验证
func (gcc *GmCryptoClient) VerifyXuperSignature(publicKeys []*ecdsa.PublicKey, sig []byte, message []byte) (valid bool, err error) {
	return signature.XuperSigVerify(publicKeys, sig, message)
}

// --- XuperSignature 统一签名相关 end ---

// --- 	hierarchical deterministic 分层确定性算法相关 start ---

// 通过助记词恢复出分层确定性根密钥
func (gcc *GmCryptoClient) GenerateMasterKeyByMnemonic(mnemonic string, language int) (string, error) {
	return hd.GenerateMasterKeyByMnemonic(mnemonic, language)
}

// 通过分层确定性私钥/公钥（如根私钥）推导出子私钥/公钥
func (gcc *GmCryptoClient) GenerateChildKey(parentKey string, index uint32) (string, error) {
	return hd.GenerateChildKey(parentKey, index)
}

// 将分层确定性私钥转化为公钥
func (gcc *GmCryptoClient) ConvertPrvKeyToPubKey(privateKey string) (string, error) {
	return hd.ConvertPrvKeyToPubKey(privateKey)
}

// 使用子公钥加密
func (gcc *GmCryptoClient) EncryptByHdKey(publicKey, msg string) (string, error) {
	return hd.Encrypt(publicKey, msg)
}

// 使用子公钥和祖先私钥（可以是推导出该子公钥的任何一级祖先私钥）解密
func (gcc *GmCryptoClient) DecryptByHdKey(publicKey, privateAncestorKey, cypherText string) (string, error) {
	return hd.Decrypt(publicKey, privateAncestorKey, cypherText)
}

// --- 	hierarchical deterministic 分层确定性算法相关 end ---

// --- secret_share 秘密分享算法相关 start ---

// 将秘密分割为碎片，totalShareNumber为碎片数量，minimumShareNumber为需要至少多少碎片才能还原出信息
func (gcc *GmCryptoClient) SecretSplit(totalShareNumber, minimumShareNumber int, secret []byte) (shares map[int]*big.Int, err error) {
	return complex_secret_share.ComplexSecretSplit(totalShareNumber, minimumShareNumber, secret)
}

// 通过收集到的碎片来还原出秘密
func (gcc *GmCryptoClient) SecretRetrieve(shares map[int]*big.Int) ([]byte, error) {
	return complex_secret_share.ComplexSecretRetrieve(shares)
}

// --- secret_share 秘密分享算法相关 end ---
