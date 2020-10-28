package xchain

/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com> 西二旗第一帅
*/

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"

	"github.com/xuperchain/crypto/client/service/base"
	"github.com/xuperchain/crypto/common/account"
	"github.com/xuperchain/crypto/core/bls_sign"
	"github.com/xuperchain/crypto/core/config"
	"github.com/xuperchain/crypto/core/ecies"
	"github.com/xuperchain/crypto/core/hash"
	"github.com/xuperchain/crypto/core/hdwallet/key"
	"github.com/xuperchain/crypto/core/multisign"
	"github.com/xuperchain/crypto/core/schnorr_ring_sign"
	"github.com/xuperchain/crypto/core/schnorr_sign"
	"github.com/xuperchain/crypto/core/secret_share/complex_secret_share"
	"github.com/xuperchain/crypto/core/sign"
	"github.com/xuperchain/crypto/core/signature"

	"github.com/xuperchain/crypto/common/zkp"
	"github.com/xuperchain/crypto/core/zkp/zk_snark/hash/mimc"

	"github.com/xuperchain/crypto/common/math/ecc"
	"github.com/xuperchain/crypto/core/threshold/schnorr/dkg"
	"github.com/xuperchain/crypto/core/threshold/schnorr/tss_sign"

	"github.com/xuperchain/crypto/common/utils"

	accountUtil "github.com/xuperchain/crypto/core/account"
	aesUtil "github.com/xuperchain/crypto/core/aes"
	hd "github.com/xuperchain/crypto/core/hdwallet/api"
	walletRand "github.com/xuperchain/crypto/core/hdwallet/rand"

	backend_bn256 "github.com/consensys/gnark/backend/bn256"
	groth16_bn256 "github.com/consensys/gnark/backend/bn256/groth16"
)

type XchainCryptoClient struct {
	base.CryptoClient
}

// --- 哈希算法相关 start ---

// 使用SHA256做单次哈希运算
func (xcc *XchainCryptoClient) HashUsingSha256(data []byte) []byte {
	hashResult := hash.HashUsingSha256(data)
	return hashResult
}

// 使用SHA256做双次哈希运算，担心SHA256存在后门时可以这么做
func (xcc *XchainCryptoClient) HashUsingDoubleSha256(data []byte) []byte {
	hashResult := hash.DoubleSha256(data)
	return hashResult
}

// 使用Hmac512做哈希运算
func (xcc *XchainCryptoClient) HashUsingHmac512(data, key []byte) []byte {
	hashResult := hash.HashUsingHmac512(data, key)
	return hashResult
}

// 使用Ripemd160做哈希运算
func (xcc *XchainCryptoClient) HashUsingRipemd160(data []byte) []byte {
	hashResult := hash.HashUsingRipemd160(data)
	return hashResult
}

// 使用MiMC做哈希运算
func (xcc *XchainCryptoClient) HashUsingDefaultMiMC(data []byte) []byte {
	hashResult := hash.HashUsingDefaultMiMC(data)
	return hashResult
}

// --- 哈希算法相关 end ---

// --- 随机数相关 start ---

// 产生随机熵
func (xcc *XchainCryptoClient) GenerateEntropy(bitSize int) ([]byte, error) {
	entropyByte, err := walletRand.GenerateEntropy(bitSize)
	return entropyByte, err
}

// --- 随机数相关 end ---

// --- 助记词相关 start ---

// 将随机熵转为助记词
func (xcc *XchainCryptoClient) GenerateMnemonic(entropy []byte, language int) (string, error) {
	mnemonic, err := walletRand.GenerateMnemonic(entropy, language)
	return mnemonic, err
}

// 将助记词转为指定长度的随机数种子，在此过程中，校验助记词是否合法
func (xcc *XchainCryptoClient) GenerateSeedWithErrorChecking(mnemonic string, password string, keyLen int, language int) ([]byte, error) {
	seed, err := walletRand.GenerateSeedWithErrorChecking(mnemonic, password, keyLen, language)
	return seed, err
}

// --- 助记词相关 end ---

// --- 密钥字符串转换相关 start ---

// 获取ECC私钥的json格式的表达的字符串
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyJsonFormatStr(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPrivateKeyJsonFormat, err := accountUtil.GetEcdsaPrivateKeyJsonFormat(k)
	return jsonEcdsaPrivateKeyJsonFormat, err
}

// 通过私钥获取ECC公钥的json格式的表达的字符串
func (xcc *XchainCryptoClient) GetEcdsaPublicKeyJsonFormatStr(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPublicKeyJsonFormat, err := accountUtil.GetEcdsaPublicKeyJsonFormat(k)
	return jsonEcdsaPublicKeyJsonFormat, err
}

// 通过公钥获取ECC公钥的json格式的表达的字符串
func (xcc *XchainCryptoClient) GetEcdsaPublicKeyJsonFormatStrFromPublicKey(k *ecdsa.PublicKey) (string, error) {
	jsonEcdsaPublicKeyJsonFormat, err := accountUtil.GetEcdsaPublicKeyJsonFormatFromPublicKey(k)
	return jsonEcdsaPublicKeyJsonFormat, err
}

// 从json格式私钥内容字符串产生ECC私钥
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyFromJsonStr(keyStr string) (*ecdsa.PrivateKey, error) {
	jsonBytes := []byte(keyStr)
	return accountUtil.GetEcdsaPrivateKeyFromJson(jsonBytes)
}

// 从json格式公钥内容字符串产生ECC公钥
func (xcc *XchainCryptoClient) GetEcdsaPublicKeyFromJsonStr(keyStr string) (*ecdsa.PublicKey, error) {
	jsonBytes := []byte(keyStr)
	return accountUtil.GetEcdsaPublicKeyFromJson(jsonBytes)
}

// --- 密钥字符串转换相关 end ---

// --- 地址生成相关 start ---

// 使用单个公钥来生成钱包地址
func (xcc *XchainCryptoClient) GetAddressFromPublicKey(key *ecdsa.PublicKey) (string, error) {
	address, err := accountUtil.GetAddressFromPublicKey(key)
	return address, err
}

// 使用多个公钥来生成钱包地址（环签名，多重签名地址）
func (xcc *XchainCryptoClient) GetAddressFromPublicKeys(keys []*ecdsa.PublicKey) (string, error) {
	address, err := accountUtil.GetAddressFromPublicKeys(keys)
	return address, err
}

// 验证钱包地址是否是合法的格式。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *XchainCryptoClient) CheckAddressFormat(address string) (bool, uint8) {
	isValid, nVersion := accountUtil.CheckAddressFormat(address)
	return isValid, nVersion
}

// 验证钱包地址是否和指定的公钥match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *XchainCryptoClient) VerifyAddressUsingPublicKey(address string, pub *ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := accountUtil.VerifyAddressUsingPublicKey(address, pub)
	return isValid, nVersion
}

// 验证钱包地址（环签名，多重签名地址）是否和指定的公钥数组match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *XchainCryptoClient) VerifyAddressUsingPublicKeys(address string, pub []*ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := accountUtil.VerifyAddressUsingPublicKeys(address, pub)
	return isValid, nVersion
}

// --- 地址生成相关 end ---

// --- 账户相关 start ---

// 通过随机数种子来生成椭圆曲线加密所需要的公钥和私钥
func (xcc *XchainCryptoClient) GenerateKeyBySeed(seed []byte) (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256()
	privateKey, err := sign.GenerateKeyBySeed(curve, seed)
	return privateKey, err
}

// ExportNewAccount 创建新账户(不使用助记词，不推荐使用)
func (xcc *XchainCryptoClient) ExportNewAccount(path string) error {
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}
	return accountUtil.ExportNewAccount(path, privateKey)
}

// 创建含有助记词的新的账户，返回的字段：（助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (xcc *XchainCryptoClient) CreateNewAccountWithMnemonic(language int, strength uint8) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := accountUtil.CreateNewAccountWithMnemonic(language, strength, config.Nist)
	return ecdsaAccount, err
}

// 创建新的账户，并用支付密码加密私钥后存在本地，
// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (xcc *XchainCryptoClient) CreateNewAccountAndSaveSecretKey(path string, language int, strength uint8, password string) (*account.ECDSAInfo, error) {
	ecdasaInfo, err := key.CreateAndSaveSecretKey(path, walletRand.SimplifiedChinese, accountUtil.StrengthHard, password, config.Nist)
	return ecdasaInfo, err
}

// 创建新的账户，并导出相关文件（含助记词）到本地。生成如下几个文件：1.助记词，2.私钥，3.公钥，4.钱包地址
func (xcc *XchainCryptoClient) ExportNewAccountWithMnemonic(path string, language int, strength uint8) error {
	err := accountUtil.ExportNewAccountWithMnemonic(path, language, strength, config.Nist)
	return err
}

// 从助记词恢复钱包账户
// TODO: 后续可以从助记词中识别出语言类型
func (xcc *XchainCryptoClient) RetrieveAccountByMnemonic(mnemonic string, language int) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := accountUtil.GenerateAccountByMnemonic(mnemonic, language)
	return ecdsaAccount, err
}

// 从助记词恢复钱包账户，并用支付密码加密私钥后存在本地，
// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (xcc *XchainCryptoClient) RetrieveAccountByMnemonicAndSavePrivKey(path string, language int, mnemonic string, password string) (*account.ECDSAInfo, error) {
	ecdsaAccount, err := key.CreateAndSaveSecretKeyWithMnemonic(path, language, mnemonic, password)
	return ecdsaAccount, err
}

// 使用支付密码加密账户信息并返回加密后的数据（后续用来回传至云端）
func (xcc *XchainCryptoClient) EncryptAccount(info *account.ECDSAAccount, password string) (*account.ECDSAAccountToCloud, error) {
	ecdsaAccountToCloud, err := key.EncryptAccount(info, password)
	return ecdsaAccountToCloud, err
}

// 从导出的私钥文件读取私钥的byte格式
func (xcc *XchainCryptoClient) GetBinaryEcdsaPrivateKeyFromFile(path string, password string) ([]byte, error) {
	binaryEcdsaPrivateKey, err := key.GetBinaryEcdsaPrivateKeyFromFile(path, password)
	return binaryEcdsaPrivateKey, err
}

// 使用支付密码从导出的私钥文件读取私钥
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyFromFileByPassword(path string, password string) (*ecdsa.PrivateKey, error) {
	ecdsaPrivateKey, err := key.GetEcdsaPrivateKeyFromFile(path, password)
	return ecdsaPrivateKey, err
}

// 使用支付密码从二进制加密字符串获取真实私钥的字节数组
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyBytesFromEncryptedStringByPassword(encryptedPrivateKey string, password string) ([]byte, error) {
	binaryEcdsaPrivateKey, err := key.GetBinaryEcdsaPrivateKeyFromString(encryptedPrivateKey, password)
	return binaryEcdsaPrivateKey, err
}

// 使用支付密码从二进制加密字符串获取真实ECC私钥
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyFromEncryptedStringByPassword(encryptedPrivateKey string, password string) (*ecdsa.PrivateKey, error) {
	binaryEcdsaPrivateKey, err := key.GetEcdsaPrivateKeyFromString(encryptedPrivateKey, password)
	return binaryEcdsaPrivateKey, err
}

// 从导出的私钥文件读取私钥
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	ecdsaPrivateKey, err := accountUtil.GetEcdsaPrivateKeyFromFile(filename)
	return ecdsaPrivateKey, err
}

// 从导出的公钥文件读取公钥
func (xcc *XchainCryptoClient) GetEcdsaPublicKeyFromFile(filename string) (*ecdsa.PublicKey, error) {
	ecdsaPublicKey, err := accountUtil.GetEcdsaPublicKeyFromFile(filename)
	return ecdsaPublicKey, err
}

// 切分账户私钥
func (xcc *XchainCryptoClient) SplitPrivateKey(jsonPrivateKey string, totalShareNumber, minimumShareNumber int) ([]string, error) {
	jsonPrivateKeyShares, err := accountUtil.SplitPrivateKey(jsonPrivateKey, totalShareNumber, minimumShareNumber)
	return jsonPrivateKeyShares, err
}

// 通过私钥片段恢复私钥
func (xcc *XchainCryptoClient) RetrievePrivateKeyByShares(jsonPrivateKeyShares []string) (string, error) {
	jsonPrivateKey, err := accountUtil.RetrievePrivateKeyByShares(jsonPrivateKeyShares)
	return jsonPrivateKey, err
}

// 将私钥的曲线转化为secp256k1，并重新计算包含的公钥
func (xcc *XchainCryptoClient) ChangePrivCurveToS256k1(key *ecdsa.PrivateKey) *ecdsa.PrivateKey {
	newPriv := utils.ChangePrivCurveToS256k1(key)
	return newPriv
}

// --- 账户相关 end ---

// --- 普通单签名相关 start ---

// 使用ECC私钥来签名
func (xcc *XchainCryptoClient) SignECDSA(k *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	signature, err := sign.SignECDSA(k, msg)
	return signature, err
}

// 使用ECC私钥来签名，生成统一签名的新签名函数
func (xcc *XchainCryptoClient) SignV2ECDSA(k *ecdsa.PrivateKey, msg []byte) ([]byte, error) {
	signature, err := sign.SignV2ECDSA(k, msg)
	return signature, err
}

// 使用ECC公钥来验证签名 -- 对应SignECDSA
func (xcc *XchainCryptoClient) VerifyECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyECDSA(k, signature, msg)
	return result, err
}

// 使用ECC公钥来验证签名，验证统一签名的新签名函数  -- 内部函数，供统一验签函数调用
func (xcc *XchainCryptoClient) VerifyV2ECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyV2ECDSA(k, signature, msg)
	return result, err
}

// --- 普通单签名相关 end ---

// --- 加解密相关 start ---

// 使用椭圆曲线非对称加密
func (xcc *XchainCryptoClient) EncryptByEcdsaKey(publicKey *ecdsa.PublicKey, msg []byte) (cypherText []byte, err error) {
	cypherText, err = ecies.Encrypt(publicKey, msg)
	return cypherText, err
}

// 使用椭圆曲线非对称解密
func (xcc *XchainCryptoClient) DecryptByEcdsaKey(privateKey *ecdsa.PrivateKey, cypherText []byte) (msg []byte, err error) {
	msg, err = ecies.Decrypt(privateKey, cypherText)
	return msg, err
}

// 使用AES对称加密算法加密
func (xcc *XchainCryptoClient) EncryptByAESKey(info string, cypherKey string) (string, error) {
	cipherInfo, err := aesUtil.Encrypt([]byte(info), []byte(cypherKey))
	if err != nil {
		return "", err
	}

	return string(cipherInfo), err
}

// 使用AES对称加密算法解密
func (xcc *XchainCryptoClient) DecryptByAESKey(cipherInfo string, cypherKey string) (string, error) {
	info, err := aesUtil.Decrypt([]byte(cipherInfo), []byte(cypherKey))
	if err != nil {
		return "", err
	}

	return string(info), nil
}

// 使用AES对称加密算法加密，密钥会被增强拓展，提升破解难度
func (xcc *XchainCryptoClient) EncryptHardenByAESKey(info string, cypherKey string) (string, error) {
	return key.EncryptByKey(info, cypherKey)
}

// 使用AES对称加密算法解密，密钥曾经被增强拓展，提升破解难度
func (xcc *XchainCryptoClient) DecryptHardenByAESKey(cipherInfo string, cypherKey string) (string, error) {
	return key.DecryptByKey(cipherInfo, cypherKey)
}

// 将经过支付密码加密的账户保存到文件中
func (xcc *XchainCryptoClient) SaveEncryptedAccountToFile(account *account.ECDSAAccountToCloud, path string) error {
	return key.SaveAccountFile(account, path)
}

// --- 加解密相关 end ---

// --- 多重签名相关 start ---

// 每个多重签名算法流程的参与节点生成32位长度的随机byte，返回值可以认为是k
func (xcc *XchainCryptoClient) GetRandom32Bytes() ([]byte, error) {
	return multisign.GetRandom32Bytes()
}

// 每个多重签名算法流程的参与节点生成Ri = Ki*G
func (xcc *XchainCryptoClient) GetRiUsingRandomBytes(key *ecdsa.PublicKey, k []byte) []byte {
	return multisign.GetRiUsingRandomBytes(key, k)
}

// 负责计算多重签名的节点来收集所有节点的Ri，并计算R = k1*G + k2*G + ... + kn*G
func (xcc *XchainCryptoClient) GetRUsingAllRi(key *ecdsa.PublicKey, arrayOfRi [][]byte) []byte {
	return multisign.GetRUsingAllRi(key, arrayOfRi)
}

// 负责计算多重签名的节点来收集所有节点的公钥Pi，并计算公共公钥：C = P1 + P2 + ... + Pn
func (xcc *XchainCryptoClient) GetSharedPublicKeyForPublicKeys(keys []*ecdsa.PublicKey) ([]byte, error) {
	return multisign.GetSharedPublicKeyForPublicKeys(keys)
}

// 负责计算多重签名的节点将计算出的R和C分别传递给各个参与节点后，由各个参与节点再次计算自己的Si
// 计算 Si = Ki + HASH(C,R,m) * Xi
// X代表大数D，也就是私钥的关键参数
func (xcc *XchainCryptoClient) GetSiUsingKCRM(key *ecdsa.PrivateKey, k []byte, c []byte, r []byte, message []byte) []byte {
	return multisign.GetSiUsingKCRM(key, k, c, r, message)
}

// 负责计算多重签名的节点来收集所有节点的Si，并计算出S = sum(si)
func (xcc *XchainCryptoClient) GetSUsingAllSi(arrayOfSi [][]byte) []byte {
	return multisign.GetSUsingAllSi(arrayOfSi)
}

// 负责计算多重签名的节点，最终生成多重签名的统一签名格式XuperSignature
func (xcc *XchainCryptoClient) GenerateMultiSignSignature(s []byte, r []byte) ([]byte, error) {
	return multisign.GenerateMultiSignSignature(s, r)
}

// 使用ECC公钥数组来进行多重签名的验证  -- 内部函数，供统一验签函数调用
func (xcc *XchainCryptoClient) VerifyMultiSig(keys []*ecdsa.PublicKey, signature, message []byte) (bool, error) {
	return multisign.VerifyMultiSig(keys, signature, message)
}

// -- 多重签名的另一种用法，适用于完全中心化的流程
// 使用ECC私钥数组来进行多重签名，生成统一签名格式XuperSignature
func (xcc *XchainCryptoClient) MultiSign(keys []*ecdsa.PrivateKey, message []byte) ([]byte, error) {
	return multisign.MultiSign(keys, message)
}

// --- 多重签名相关 end ---

// --- schnorr签名算法相关 start ---

// schnorr签名算法 生成统一签名XuperSignature
func (xcc *XchainCryptoClient) SignSchnorr(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	return schnorr_sign.Sign(privateKey, message)
}

// schnorr签名算法 验证签名  -- 内部函数，供统一验签函数调用
func (xcc *XchainCryptoClient) VerifySchnorr(publicKey *ecdsa.PublicKey, sig, message []byte) (bool, error) {
	return schnorr_sign.Verify(publicKey, sig, message)
}

// --- schnorr签名算法相关 end ---

// --- schnorr 环签名算法相关 start ---

// schnorr环签名算法 生成统一签名XuperSignature
func (xcc *XchainCryptoClient) SignSchnorrRing(keys []*ecdsa.PublicKey, privateKey *ecdsa.PrivateKey, message []byte) (ringSignature []byte, err error) {
	return schnorr_ring_sign.Sign(keys, privateKey, message)
}

// schnorr环签名算法 验证签名  -- 内部函数，供统一验签函数调用
func (xcc *XchainCryptoClient) VerifySchnorrRing(keys []*ecdsa.PublicKey, sig, message []byte) (bool, error) {
	return schnorr_ring_sign.Verify(keys, sig, message)
}

// --- schnorr 环签名算法相关 end ---

// --- hierarchical deterministic 分层确定性算法相关 start ---

// 通过助记词恢复出分层确定性根密钥
func (xcc *XchainCryptoClient) GenerateMasterKeyByMnemonic(mnemonic string, language int) (string, error) {
	return hd.GenerateMasterKeyByMnemonic(mnemonic, language)
}

// 通过分层确定性私钥/公钥（如根私钥）推导出子私钥/公钥
func (xcc *XchainCryptoClient) GenerateChildKey(parentKey string, index uint32) (string, error) {
	return hd.GenerateChildKey(parentKey, index)
}

// 将分层确定性私钥转化为公钥
func (xcc *XchainCryptoClient) ConvertPrvKeyToPubKey(privateKey string) (string, error) {
	return hd.ConvertPrvKeyToPubKey(privateKey)
}

// 使用子公钥加密
func (xcc *XchainCryptoClient) EncryptByHdKey(publicKey, msg string) (string, error) {
	return hd.Encrypt(publicKey, msg)
}

// 使用子公钥和祖先私钥（可以是推导出该子公钥的任何一级祖先私钥）解密
func (xcc *XchainCryptoClient) DecryptByHdKey(publicKey, privateAncestorKey, cypherText string) (string, error) {
	return hd.Decrypt(publicKey, privateAncestorKey, cypherText)
}

// --- hierarchical deterministic 分层确定性算法相关 end ---

// --- secret_share 秘密分享算法相关 start ---

// 将秘密分割为碎片，totalShareNumber为碎片数量，minimumShareNumber为需要至少多少碎片才能还原出信息
func (xcc *XchainCryptoClient) SecretSplit(totalShareNumber, minimumShareNumber int, secret []byte) (shares map[int]*big.Int, err error) {
	curve := elliptic.P256()
	return complex_secret_share.ComplexSecretSplit(totalShareNumber, minimumShareNumber, secret, curve)
}

// 通过收集到的碎片来还原出秘密
func (xcc *XchainCryptoClient) SecretRetrieve(shares map[int]*big.Int) ([]byte, error) {
	curve := elliptic.P256()
	return complex_secret_share.ComplexSecretRetrieve(shares, curve)
}

// --- secret_share 秘密分享算法相关 end ---

// --- 零知识证明算法相关 start ---

// 初始化哈希算法MiMC的参数
func (xcc *XchainCryptoClient) ZkpSetupMiMC() *zkp.ZkpInfo {
	return mimc.Setup()
}

func (xcc *XchainCryptoClient) ZkpProveMiMC(r1cs *backend_bn256.R1CS, pk *groth16_bn256.ProvingKey, secret []byte) (*groth16_bn256.Proof, error) {
	return mimc.Prove(r1cs, pk, secret)
}

func (xcc *XchainCryptoClient) ZkpVerifyMiMC(proof *groth16_bn256.Proof, vk *groth16_bn256.VerifyingKey, hashResult []byte) (bool, error) {
	return mimc.Verify(proof, vk, hashResult)
}

// --- 零知识证明算法相关 end ---

// --- 门限签名相关 start ---

// -- STEP 1: DKG - distributed key generation --
// distributed private key generation and distributed public key generation

// - method 1 start -
// 一个步骤整体
// 所有潜在参与节点根据门限目标生成产生本地秘密和验证点的私钥碎片
// minimumShareNumber可以理解为threshold，至少需要minimumShareNumber个潜在参与节点进行实际参与才能完成门限签名
func (xcc *XchainCryptoClient) GetLocalShares(totalShareNumber, minimumShareNumber int) (shares map[int]*big.Int, points []*ecc.Point, err error) {
	return dkg.LocalSecretShareGenerateWithVerifyPoints(totalShareNumber, minimumShareNumber)
}

// - method 1 end -

// - method 2 start -
// 分步骤
// 为产生本地秘密的私钥碎片做准备，预先生成好一个目标多项式
// minimumShareNumber可以理解为threshold，至少需要minimumShareNumber个潜在参与节点进行实际参与才能完成门限签名
func (xcc *XchainCryptoClient) GetPolynomialForSecretShareGenerate(totalShareNumber, minimumShareNumber int) ([]*big.Int, error) {
	return dkg.GetPolynomialForSecretShareGenerate(totalShareNumber, minimumShareNumber)
}

// 为产生本地秘密的私钥碎片做准备，通过目标多项式生成验证点
func (xcc *XchainCryptoClient) GetVerifyPointByPolynomial(poly []*big.Int) (*ecc.Point, error) {
	return dkg.GetVerifyPointByPolynomial(poly)
}

// 为产生本地秘密的私钥碎片做准备，通过目标多项式和节点index生成对应的碎片
func (xcc *XchainCryptoClient) GetSpecifiedSecretShareByPolynomial(poly []*big.Int, index *big.Int) *big.Int {
	return dkg.GetSpecifiedSecretShareByPolynomial(poly, index)
}

// - method 2 end -

// ---
// TODO: 网络层，各个潜在参与节点将对应的密钥碎片和验证点数据发送给对应的其它潜在参与节点
// shares的key就是目标对应节点的index
// 每个潜在参与节点都有一个unique的index作为标志
// 网络层在其他的模块进行实现
// ---

// 每个潜在参与节点根据所收集的所有的与自己相关的碎片(自己的Index是X值，收集所有该X值对应的Y值)，
// 来计算出自己的本地私钥X(i)(该X值对应的Y值之和)，这是一个关键秘密信息
func (xcc *XchainCryptoClient) GetLocalPrivateKeyByShares(shares []*big.Int) *ecdsa.PrivateKey {
	return dkg.LocalPrivateKeyGenerate(shares)
}

// 每个潜在参与节点来收集所有节点的秘密验证点，并计算公共公钥：C = VP(1) + VP(2) + ... + VP(i)
func (xcc *XchainCryptoClient) GetSharedPublicKey(verifyPoints []*ecc.Point) (*ecdsa.PublicKey, error) {
	return dkg.PublicKeyGenerate(verifyPoints)
}

// TODO：存储层，各个节点保留相关信息：自己的私钥，自己的编号Index

// -- STEP 2: DSG - distributed signature generation --

// 每个实际参与节点生成32位长度的随机byte，返回值可以认为是K(i)
// 可以跟多重签名复用同名函数
//func (xcc *XchainCryptoClient) GetRandom32Bytes() ([]byte, error) {
//	return tss_sign.GetRandom32Bytes()
//}

// 每个门限签名算法流程的实际参与节点生成R(i) = K(i)*G
// 然后将自己的R(i)广播到全网其它节点
// 可以跟多重签名复用同名函数
//func (xcc *XchainCryptoClient) GetRiUsingRandomBytes(key *ecdsa.PublicKey, k []byte) []byte {
//	return tss_sign.GetRiUsingRandomBytes(key, k)
//}

// 每个实际参与节点都收集所有实际参与节点的R(i)，
// 并计算R = sum(R(i)) = K(1)*G + K(2)*G + ... + K(i)*G
// 可以跟多重签名复用同名函数
//func (xcc *XchainCryptoClient) GetRUsingAllRi(key *ecdsa.PublicKey, arrayOfRi [][]byte) []byte {
//	return tss_sign.GetRUsingAllRi(key, arrayOfRi)
//}

// 每个实际参与节点再次计算自己的独有系数与自己私钥秘密的乘积，也就是X(i) * Coef(i)，为下一步的S(i)计算做准备
// indexSet是指所有实际参与节点的index所组成的集合
// localIndexPos是本节点在indexSet中的位置
// key是在DKG过程中，自己计算出的私钥
func (xcc *XchainCryptoClient) GetXiWithcoef(indexSet []*big.Int, localIndexPos int, key *ecdsa.PrivateKey) *big.Int {
	return tss_sign.GetXiWithcoef(indexSet, localIndexPos, key)
}

// 每个实际参与节点再次计算自己的S(i)
// S(i) = K(i) + HASH(C,R,m) * X(i) * Coef(i)
// X代表大数D，也就是私钥的关键参数
func (xcc *XchainCryptoClient) GetSiUsingKCRMWithCoef(k []byte, c []byte, r []byte, message []byte, coef *big.Int) []byte {
	return tss_sign.GetSiUsingKCRMWithCoef(k, c, r, message, coef)
}

// 注意：专用于多层门限算法，每个实际参与节点再次计算自己的S(i) 版本2
// S(i) = HASH(C,R,m) * X(i) * Coef(i)
// X代表大数D，也就是私钥的关键参数
func (xcc *XchainCryptoClient) GetSiUsingKCRMWithCoefNoKi(c []byte, r []byte, message []byte, coef *big.Int) []byte {
	return tss_sign.GetSiUsingKCRMWithCoefNoKi(c, r, message, coef)
}

// 负责计算门限签名的实际参与节点来收集所有实际参与节点的S(i)，并计算出S = sum(S(i))
// = K(1) + HASH(C,R,m) * X(1) * Coef(1) + K(2) + HASH(C,R,m) * X(2) * Coef(2) + ... + K(i) + HASH(C,R,m) * X(i) * Coef(i)
// 可以跟多重签名复用同名函数
//func (xcc *XchainCryptoClient) GetSUsingAllSi(arrayOfSi [][]byte) []byte {
//	return tss_sign.GetSUsingAllSi(arrayOfSi)
//}

// 负责计算门限签名的节点，最终生成门限签名的统一签名格式XuperSignature
func (xcc *XchainCryptoClient) GenerateTssSignSignature(s []byte, r []byte) ([]byte, error) {
	return tss_sign.GenerateTssSignSignature(s, r)
}

// 使用ECC公钥来进行门限签名的验证  -- 内部函数，供统一验签函数调用
func (xcc *XchainCryptoClient) VerifyTssSig(key *ecdsa.PublicKey, signature, message []byte) (bool, error) {
	return tss_sign.VerifyTssSig(key, signature, message)
}

// --- 门限签名相关 end ---

// --- BLS签名相关 start ---

// BLS签名算法 生成公钥和私钥对
func (xcc *XchainCryptoClient) GenerateBlsKeyPair() (*bls_sign.PrivateKey, *bls_sign.PublicKey) {
	return bls_sign.GenerateKeyPair()
}

// BLS签名算法 生成统一签名XuperSignature
func (xcc *XchainCryptoClient) SignBls(privateKey *bls_sign.PrivateKey, message []byte) (blsSignature []byte, err error) {
	return bls_sign.Sign(privateKey, message)
}

// 使用BLS公钥来进行门限签名的验证  -- 外部函数，因为椭圆曲线的原因，暂时无法成为内部函数，供统一验签函数调用
func (xcc *XchainCryptoClient) VerifyBlsSig(key *bls_sign.PublicKey, signature, message []byte) (bool, error) {
	return bls_sign.Verify(key, signature, message)
}

// --- BLS签名相关 end ---

// --- XuperSignature 统一签名相关 start ---

// --- 统一验签算法，可以对用各种签名算法生成的统一签名格式XuperSignature进行验证
func (xcc *XchainCryptoClient) VerifyXuperSignature(publicKeys []*ecdsa.PublicKey, sig []byte, message []byte) (valid bool, err error) {
	return signature.XuperSigVerify(publicKeys, sig, message)
}

// --- XuperSignature 统一签名相关 end ---
