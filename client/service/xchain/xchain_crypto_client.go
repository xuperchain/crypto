package xchain

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"

	"github.com/xuperchain/crypto/core/account"
	"github.com/xuperchain/crypto/core/config"
	"github.com/xuperchain/crypto/core/ecies"
	"github.com/xuperchain/crypto/core/hash"
	"github.com/xuperchain/crypto/core/hdwallet/key"
	"github.com/xuperchain/crypto/core/multisign"
	"github.com/xuperchain/crypto/core/schnorr_ring_sign"
	"github.com/xuperchain/crypto/core/schnorr_sign"
	"github.com/xuperchain/crypto/core/sign"
	"github.com/xuperchain/crypto/core/signature"

	aesUtil "github.com/xuperchain/crypto/core/aes"
	hd "github.com/xuperchain/crypto/core/hdwallet/api"
	walletRand "github.com/xuperchain/crypto/core/hdwallet/rand"
)

type XchainCryptoClient struct {
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

// --- 哈希算法相关 end ---

// 产生随机熵
func (xcc *XchainCryptoClient) GenerateEntropy(bitSize int) ([]byte, error) {
	entropyByte, err := walletRand.GenerateEntropy(bitSize)
	return entropyByte, err
}

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

// 获取ECC私钥的json格式的表达
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPrivateKeyJsonFormat, err := account.GetEcdsaPrivateKeyJsonFormat(k)
	return jsonEcdsaPrivateKeyJsonFormat, err
}

// 获取ECC公钥的json格式的表达
func (xcc *XchainCryptoClient) GetEcdsaPublicKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	jsonEcdsaPublicKeyJsonFormat, err := account.GetEcdsaPublicKeyJsonFormat(k)
	return jsonEcdsaPublicKeyJsonFormat, err
}

// --- 地址生成相关 start ---

// 使用单个公钥来生成钱包地址
func (xcc *XchainCryptoClient) GetAddressFromPublicKey(key *ecdsa.PublicKey) (string, error) {
	address, err := account.GetAddressFromPublicKey(key)
	return address, err
}

// 使用多个公钥来生成钱包地址
func (xcc *XchainCryptoClient) GetAddressFromPublicKeys(keys []*ecdsa.PublicKey) (string, error) {
	address, err := account.GetAddressFromPublicKeys(keys)
	return address, err
}

// 验证钱包地址是否是合法的格式。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *XchainCryptoClient) CheckAddressFormat(address string) (bool, uint8) {
	isValid, nVersion := account.CheckAddressFormat(address)
	return isValid, nVersion
}

// 验证钱包地址是否和指定的公钥match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *XchainCryptoClient) VerifyAddressUsingPublicKey(address string, pub *ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := account.VerifyAddressUsingPublicKey(address, pub)
	return isValid, nVersion
}

// 验证钱包地址是否和指定的公钥数组match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
func (xcc *XchainCryptoClient) VerifyAddressUsingPublicKeys(address string, pub []*ecdsa.PublicKey) (bool, uint8) {
	isValid, nVersion := account.VerifyAddressUsingPublicKeys(address, pub)
	return isValid, nVersion
}

// --- 地址生成相关 end ---

// 通过随机数种子来生成椭圆曲线加密所需要的公钥和私钥
func (xcc *XchainCryptoClient) GenerateKeyBySeed(seed []byte) (*ecdsa.PrivateKey, error) {
	curve := elliptic.P256()
	privateKey, err := sign.GenerateKeyBySeed(curve, seed)
	return privateKey, err
}

//// 获取ECC私钥的json格式的表达
//func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
//	jsonEcdsaPrivateKeyJsonFormat, err := account.GetEcdsaPrivateKeyJsonFormat(k)
//	return jsonEcdsaPrivateKeyJsonFormat, err
//}
//
//// 获取ECC公钥的json格式的表达
//func (xcc *XchainCryptoClient) GetEcdsaPublicKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
//	jsonEcdsaPublicKeyJsonFormat, err := account.GetEcdsaPublicKeyJsonFormat(k)
//	return jsonEcdsaPublicKeyJsonFormat, err
//}

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

// 使用ECC公钥来验证签名，验证统一签名的新签名函数
func (xcc *XchainCryptoClient) VerifyECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyECDSA(k, signature, msg)
	return result, err
}

// 使用ECC公钥来验证签名，验证统一签名的新签名函数
func (xcc *XchainCryptoClient) VerifyV2ECDSA(k *ecdsa.PublicKey, signature, msg []byte) (bool, error) {
	result, err := sign.VerifyV2ECDSA(k, signature, msg)
	return result, err
}

//// 使用公钥来生成钱包地址
//func (xcc *XchainCryptoClient) GetAddressFromPublicKey(nVersion uint8, pub *ecdsa.PublicKey) string {
//	address := account.GetAddressFromPublicKey(nVersion, pub)
//	return address
//}

// ExportNewAccount 创建新账户(不使用助记词，不推荐使用)
func (xcc *XchainCryptoClient) ExportNewAccount(path string) error {
	lowLevelPrivateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return err
	}
	return account.ExportNewAccount(path, lowLevelPrivateKey)
}

// 创建含有助记词的新的账户，返回的字段：（助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (xcc *XchainCryptoClient) CreateNewAccountWithMnemonic(language int, strength uint8) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := account.CreateNewAccountWithMnemonic(language, strength, config.Nist)
	return ecdsaAccount, err
}

// 创建新的账户，并用支付密码加密私钥后存在本地，
// 返回的字段：（随机熵（供其他钱包软件推导出私钥）、助记词、私钥的json、公钥的json、钱包地址） as ECDSAAccount，以及可能的错误信息
func (xcc *XchainCryptoClient) CreateNewAccountAndSaveSecretKey(path string, language int, strength uint8, password string) (*account.ECDSAInfo, error) {
	ecdasaInfo, err := key.CreateAndSaveSecretKey(path, walletRand.SimplifiedChinese, account.StrengthHard, password, config.Nist)
	return ecdasaInfo, err
}

// 创建新的账户，并导出相关文件（含助记词）到本地。生成如下几个文件：1.助记词，2.私钥，3.公钥，4.钱包地址
func (xcc *XchainCryptoClient) ExportNewAccountWithMnemonic(path string, language int, strength uint8) error {
	err := account.ExportNewAccountWithMnemonic(path, language, strength, config.Nist)
	return err
}

// 从助记词恢复钱包账户
// TODO: 后续可以从助记词中识别出语言类型
func (xcc *XchainCryptoClient) RetrieveAccountByMnemonic(mnemonic string, language int) (*account.ECDSAAccount, error) {
	ecdsaAccount, err := account.GenerateAccountByMnemonic(mnemonic, language)
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

// 从二进制加密字符串获取真实私钥的byte格式
func (xcc *XchainCryptoClient) GetBinaryEcdsaPrivateKeyFromString(encryptPrivateKey string, password string) ([]byte, error) {
	binaryEcdsaPrivateKey, err := key.GetBinaryEcdsaPrivateKeyFromString(encryptPrivateKey, password)
	return binaryEcdsaPrivateKey, err
}

// 从导出的私钥文件读取私钥
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyFromFile(filename string) (*ecdsa.PrivateKey, error) {
	ecdsaPrivateKey, err := account.GetEcdsaPrivateKeyFromFile(filename)
	return ecdsaPrivateKey, err
}

// 从导出的公钥文件读取公钥
func (xcc *XchainCryptoClient) GetEcdsaPublicKeyFromFile(filename string) (*ecdsa.PublicKey, error) {
	ecdsaPublicKey, err := account.GetEcdsaPublicKeyFromFile(filename)
	return ecdsaPublicKey, err
}

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

// 从导出的私钥文件读取私钥
func (xcc *XchainCryptoClient) GetEcdsaPrivateKeyFromJson(jsonBytes []byte) (*ecdsa.PrivateKey, error) {
	return account.GetEcdsaPrivateKeyFromJson(jsonBytes)
}

// 从导出的公钥文件读取公钥
func (xcc *XchainCryptoClient) GetEcdsaPublicKeyFromJson(jsonBytes []byte) (*ecdsa.PublicKey, error) {
	return account.GetEcdsaPublicKeyFromJson(jsonBytes)
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

// 使用ECC公钥数组来进行多重签名的验证
func (xcc *XchainCryptoClient) VerifyMultiSig(keys []*ecdsa.PublicKey, signature, message []byte) (bool, error) {
	if len(keys) < 2 {
		return false, fmt.Errorf("The total num of keys should be greater than two.")
	}

	// 判断是否是nist标准的私钥
	switch keys[0].Params().Name {
	case config.CurveNist: // NIST
		signature, err := multisign.VerifyMultiSig(keys, signature, message)
		return signature, err
	case config.CurveGm: // 国密
		return false, fmt.Errorf("This cryptography has not been supported yet.")
	default: // 不支持的密码学类型
		return false, fmt.Errorf("This cryptography has not been supported yet.")
	}
}

// -- 多重签名的另一种用法，适用于完全中心化的流程
// 使用ECC私钥数组来进行多重签名，生成统一签名格式XuperSignature
func (xcc *XchainCryptoClient) MultiSign(keys []*ecdsa.PrivateKey, message []byte) ([]byte, error) {
	// 判断是否是nist标准的私钥
	if len(keys) < 2 {
		return nil, fmt.Errorf("The total num of keys should be greater than two.")
	}

	switch keys[0].Params().Name {
	case config.CurveNist: // NIST
		signature, err := multisign.MultiSign(keys, message)
		return signature, err
	case config.CurveGm: // 国密
		return nil, fmt.Errorf("This cryptography has not been supported yet.")
	default: // 不支持的密码学类型
		return nil, fmt.Errorf("This cryptography has not been supported yet.")
	}
}

// --- 多重签名相关 end ---

// --- 	schnorr签名算法相关 start ---

// schnorr签名算法 生成统一签名XuperSignature
func (xcc *XchainCryptoClient) SignSchnorr(privateKey *ecdsa.PrivateKey, message []byte) ([]byte, error) {
	return schnorr_sign.Sign(privateKey, message)
}

// schnorr签名算法 验证签名
func (xcc *XchainCryptoClient) VerifySchnorr(publicKey *ecdsa.PublicKey, sig, message []byte) (bool, error) {
	return schnorr_sign.Verify(publicKey, sig, message)
}

// --- 	schnorr签名算法相关 end ---

// --- 	schnorr 环签名算法相关 start ---

// schnorr环签名算法 生成统一签名XuperSignature
func (xcc *XchainCryptoClient) SignSchnorrRing(keys []*ecdsa.PublicKey, privateKey *ecdsa.PrivateKey, message []byte) (ringSignature []byte, err error) {
	return schnorr_ring_sign.Sign(keys, privateKey, message)
}

// schnorr环签名算法 验证签名
func (xcc *XchainCryptoClient) VerifySchnorrRing(keys []*ecdsa.PublicKey, sig, message []byte) (bool, error) {
	return schnorr_ring_sign.Verify(keys, sig, message)
}

// --- 	schnorr 环签名算法相关 end ---

// --- 统一验签算法，可以对用各种签名算法生成的签名进行验证
func (xcc *XchainCryptoClient) VerifyXuperSignature(publicKeys []*ecdsa.PublicKey, sig []byte, message []byte) (valid bool, err error) {
	return signature.XuperSigVerify(publicKeys, sig, message)
}

// --- 	hierarchical deterministic 分层确定性算法相关 start ---

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

// --- 	hierarchical deterministic 分层确定性算法相关 end ---
