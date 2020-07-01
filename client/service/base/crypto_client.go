/*
Copyright Baidu Inc. All Rights Reserved.

jingbo@baidu.com
*/

package base

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/xuperchain/crypto/common/account"
)

// interface of all Crypto functions
type CryptoClient interface {
	// --- 哈希算法相关 start ---

	//	// 使用SHA256做单次哈希运算
	//	HashUsingSM3(data []byte) []byte

	//	// 使用Hmac512做单次哈希运算
	//	HashUsingHmac512(data, key []byte) []byte

	//	// 使用Ripemd160做单次哈希运算
	//	HashUsingRipemd160(data []byte) []byte

	//	// 使用SHA256做单次哈希运算
	//	HashUsingSha256(data []byte) []byte
	//
	//	// 使用SHA256做双次哈希运算，担心SHA256存在后门时可以这么做
	//	HashUsingDoubleSha256(data []byte) []byte

	// --- 哈希算法相关 end ---

	// --- 随机数相关 start ---

	// 产生随机熵
	GenerateEntropy(bitSize int) ([]byte, error)

	// --- 随机数相关 end ---

	// --- 助记词相关 start ---

	// 将随机熵转为助记词
	GenerateMnemonic(entropy []byte, language int) (string, error)

	// 将助记词转为指定长度的随机数种子，在此过程中，校验助记词是否合法
	GenerateSeedWithErrorChecking(mnemonic string, password string, keyLen int, language int) ([]byte, error)

	// --- 助记词相关 end ---

	// --- 密钥字符串转换相关 start ---

	// 获取ECC私钥的json格式的表达
	GetEcdsaPrivateKeyJsonFormatStr(k *ecdsa.PrivateKey) (string, error)

	// 获取ECC公钥的json格式的表达
	GetEcdsaPublicKeyJsonFormatStr(k *ecdsa.PrivateKey) (string, error)

	// 从json格式私钥内容字符串产生ECC私钥
	GetEcdsaPrivateKeyFromJsonStr(keyStr string) (*ecdsa.PrivateKey, error)

	// 从json格式公钥内容字符串产生ECC公钥
	GetEcdsaPublicKeyFromJsonStr(keyStr string) (*ecdsa.PublicKey, error)

	// --- 密钥字符串转换相关 end ---

	// --- 地址生成相关 start ---

	// 使用单个公钥来生成钱包地址
	GetAddressFromPublicKey(key *ecdsa.PublicKey) (string, error)

	// 使用多个公钥来生成钱包地址（环签名，多重签名地址）
	GetAddressFromPublicKeys(keys []*ecdsa.PublicKey) (string, error)

	// 验证钱包地址是否是合法的格式。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
	CheckAddressFormat(address string) (bool, uint8)

	// 验证钱包地址是否和指定的公钥match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
	VerifyAddressUsingPublicKey(address string, pub *ecdsa.PublicKey) (bool, uint8)

	// 验证钱包地址（环签名，多重签名地址）是否和指定的公钥数组match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
	VerifyAddressUsingPublicKeys(address string, pub []*ecdsa.PublicKey) (bool, uint8)

	// --- 地址生成相关 end ---

	// --- 账户相关 start ---

	GenerateKeyBySeed(seed []byte) (*ecdsa.PrivateKey, error)

	// ExportNewAccount 创建新账户(不使用助记词，不推荐使用)
	ExportNewAccount(path string) error

	// 创建含有助记词的新的账户，返回的字段：（助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
	CreateNewAccountWithMnemonic(language int, strength uint8) (*account.ECDSAAccount, error)

	// 创建新的账户，并用支付密码加密私钥后存在本地，
	// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
	CreateNewAccountAndSaveSecretKey(path string, language int, strength uint8, password string) (*account.ECDSAInfo, error)

	// 创建新的账户，并导出相关文件（含助记词）到本地。生成如下几个文件：1.助记词，2.私钥，3.公钥，4.钱包地址
	ExportNewAccountWithMnemonic(path string, language int, strength uint8) error

	// 从助记词恢复钱包账户
	// TODO: 后续可以从助记词中识别出语言类型
	RetrieveAccountByMnemonic(mnemonic string, language int) (*account.ECDSAAccount, error)

	// 从助记词恢复钱包账户，并用支付密码加密私钥后存在本地，
	// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
	RetrieveAccountByMnemonicAndSavePrivKey(path string, language int, mnemonic string, password string) (*account.ECDSAInfo, error)

	// 使用支付密码加密账户信息并返回加密后的数据（后续用来回传至云端）
	EncryptAccount(info *account.ECDSAAccount, password string) (*account.ECDSAAccountToCloud, error)

	// 从导出的私钥文件读取私钥的byte格式
	GetBinaryEcdsaPrivateKeyFromFile(path string, password string) ([]byte, error)

	// 使用支付密码从导出的私钥文件读取私钥
	GetEcdsaPrivateKeyFromFileByPassword(path string, password string) (*ecdsa.PrivateKey, error)

	// 使用支付密码从二进制加密字符串获取真实私钥的字节数组
	GetEcdsaPrivateKeyBytesFromEncryptedStringByPassword(encryptedPrivateKey string, password string) ([]byte, error)

	// 使用支付密码从二进制加密字符串获取真实ECC私钥
	GetEcdsaPrivateKeyFromEncryptedStringByPassword(encryptedPrivateKey string, password string) (*ecdsa.PrivateKey, error)

	// 从导出的私钥文件读取私钥
	GetEcdsaPrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error)

	// 从导出的公钥文件读取公钥
	GetEcdsaPublicKeyFromFile(filename string) (*ecdsa.PublicKey, error)

	// 切分账户私钥
	SplitPrivateKey(jsonPrivateKey string, totalShareNumber, minimumShareNumber int) ([]string, error)

	// 通过私钥片段恢复私钥
	RetrievePrivateKeyByShares(jsonPrivateKeyShares []string) (string, error)

	// --- 账户相关 end ---

	// --- 普通单签名相关 start ---

	// 使用ECC私钥来签名
	SignECDSA(k *ecdsa.PrivateKey, msg []byte) ([]byte, error)

	// 使用ECC私钥来签名，生成统一签名的新签名函数
	SignV2ECDSA(k *ecdsa.PrivateKey, msg []byte) ([]byte, error)

	// 使用ECC公钥来验证签名，验证统一签名的新签名函数
	VerifyECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error)

	// 使用ECC公钥来验证签名，验证统一签名的新签名函数  -- 内部函数，供统一验签函数调用
	VerifyV2ECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error)

	// --- 普通单签名相关 end ---

	// --- 加解密相关 start ---

	// 使用椭圆曲线非对称加密
	EncryptByEcdsaKey(publicKey *ecdsa.PublicKey, msg []byte) (cypherText []byte, err error)

	// 使用椭圆曲线非对称解密
	DecryptByEcdsaKey(privateKey *ecdsa.PrivateKey, cypherText []byte) (msg []byte, err error)

	// 使用AES对称加密算法加密
	EncryptByAESKey(info string, cypherKey string) (string, error)

	// 使用AES对称加密算法解密
	DecryptByAESKey(cipherInfo string, cypherKey string) (string, error)

	// 使用AES对称加密算法加密，密钥会被增强拓展，提升破解难度
	EncryptHardenByAESKey(info string, cypherKey string) (string, error)

	// 使用AES对称加密算法解密，密钥曾经被增强拓展，提升破解难度
	DecryptHardenByAESKey(cipherInfo string, cypherKey string) (string, error)

	// 将经过支付密码加密的账户保存到文件中
	SaveEncryptedAccountToFile(account *account.ECDSAAccountToCloud, path string) error

	// --- 加解密相关 end ---

	// --- 多重签名相关 start ---

	// 每个多重签名算法流程的参与节点生成32位长度的随机byte，返回值可以认为是k
	GetRandom32Bytes() ([]byte, error)

	// 每个多重签名算法流程的参与节点生成Ri = Ki*G
	GetRiUsingRandomBytes(key *ecdsa.PublicKey, k []byte) []byte

	// 负责计算多重签名的节点来收集所有节点的Ri，并计算R = k1*G + k2*G + ... + kn*G
	GetRUsingAllRi(key *ecdsa.PublicKey, arrayOfRi [][]byte) []byte

	// 负责计算多重签名的节点来收集所有节点的公钥Pi，并计算公共公钥：C = P1 + P2 + ... + Pn
	GetSharedPublicKeyForPublicKeys(keys []*ecdsa.PublicKey) ([]byte, error)

	// 负责计算多重签名的节点将计算出的R和C分别传递给各个参与节点后，由各个参与节点再次计算自己的Si
	// 计算 Si = Ki + HASH(C,R,m) * Xi
	// X代表大数D，也就是私钥的关键参数
	GetSiUsingKCRM(key *ecdsa.PrivateKey, k []byte, c []byte, r []byte, message []byte) []byte

	// 负责计算多重签名的节点来收集所有节点的Si，并计算出S = sum(si)
	GetSUsingAllSi(arrayOfSi [][]byte) []byte

	// 负责计算多重签名的节点，最终生成多重签名的统一签名格式XuperSignature
	GenerateMultiSignSignature(s []byte, r []byte) ([]byte, error)

	// 使用ECC公钥数组来进行多重签名的验证  -- 内部函数，供统一验签函数调用
	VerifyMultiSig(keys []*ecdsa.PublicKey, signature, message []byte) (bool, error)

	// -- 多重签名的另一种用法，适用于完全中心化的流程
	// 使用ECC私钥数组来进行多重签名，生成统一签名格式XuperSignature
	MultiSign(keys []*ecdsa.PrivateKey, message []byte) ([]byte, error)

	// --- 多重签名相关 end ---

	// --- 	schnorr签名算法相关 start ---

	// schnorr签名算法 生成统一签名XuperSignature
	SignSchnorr(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error)

	// schnorr签名算法 验证签名  -- 内部函数，供统一验签函数调用
	VerifySchnorr(publicKey *ecdsa.PublicKey, sig, message []byte) (bool, error)

	// --- 	schnorr签名算法相关 end ---

	// --- 	schnorr 环签名算法相关 start ---

	// schnorr环签名算法 生成统一签名XuperSignature
	SignSchnorrRing(keys []*ecdsa.PublicKey, privateKey *ecdsa.PrivateKey, message []byte) (ringSignature []byte, err error)

	// schnorr环签名算法 验证签名  -- 内部函数，供统一验签函数调用
	VerifySchnorrRing(keys []*ecdsa.PublicKey, sig, message []byte) (bool, error)

	// --- 	schnorr 环签名算法相关 end ---

	// --- XuperSignature 统一签名相关 start ---

	// --- 统一验签算法，可以对用各种签名算法生成的签名进行验证
	VerifyXuperSignature(publicKeys []*ecdsa.PublicKey, sig []byte, message []byte) (valid bool, err error)

	// --- XuperSignature 统一签名相关 end ---

	// --- 	hierarchical deterministic 分层确定性算法相关 start ---

	// 通过助记词恢复出分层确定性根密钥
	GenerateMasterKeyByMnemonic(mnemonic string, language int) (string, error)

	// 通过分层确定性私钥/公钥（如根私钥）推导出子私钥/公钥
	GenerateChildKey(parentKey string, index uint32) (string, error)

	// 将分层确定性私钥转化为公钥
	ConvertPrvKeyToPubKey(privateKey string) (string, error)

	// 使用子公钥加密
	EncryptByHdKey(publicKey, msg string) (string, error)

	// 使用子公钥和祖先私钥（可以是推导出该子公钥的任何一级祖先私钥）解密
	DecryptByHdKey(publicKey, privateAncestorKey, cypherText string) (string, error)

	// --- 	hierarchical deterministic 分层确定性算法相关 end ---

	// --- secret_share 秘密分享算法相关 start ---

	// 将秘密分割为碎片，totalShareNumber为碎片数量，minimumShareNumber为需要至少多少碎片才能还原出信息
	SecretSplit(totalShareNumber, minimumShareNumber int, secret []byte) (shares map[int]*big.Int, err error)

	// 通过收集到的碎片来还原出秘密
	SecretRetrieve(shares map[int]*big.Int) ([]byte, error)

	// --- secret_share 秘密分享算法相关 end ---
}
