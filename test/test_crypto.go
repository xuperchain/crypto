package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"log"
	"math/big"

	"github.com/xuperchain/crypto/client/service/xchain"
	"github.com/xuperchain/crypto/core/account"
	"github.com/xuperchain/crypto/core/hdwallet/rand"
	"github.com/xuperchain/crypto/core/secret_share/complex_secret_share"

	hdapi "github.com/xuperchain/crypto/core/hdwallet/api"
)

func main() {
	xcc := new(xchain.XchainCryptoClient)

	hashResult := xcc.HashUsingSha256([]byte("This is xchain crypto"))
	log.Printf("Hash result for [This is xchain crypto] is: %s", hashResult)

	ecdsaAccount, err := xcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
	if err != nil {
		log.Printf("CreateNewAccountWithMnemonic failed and err is: %v", err)
		return
	}
	log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)

	strJsonPrivateKey := ecdsaAccount.JsonPrivateKey
	originalPrivateKey, err := xcc.GetEcdsaPrivateKeyFromJson([]byte(strJsonPrivateKey))

	// 从助记词恢复账户
	ecdsaAccount, err = xcc.RetrieveAccountByMnemonic(ecdsaAccount.Mnemonic, rand.SimplifiedChinese)
	if err != nil {
		log.Printf("RetrieveAccountByMnemonic failed and err is: %v", err)
		return
	}
	log.Printf("retrieve account from mnemonic %v, ecdsaAccount is %v and err is %v", ecdsaAccount.Mnemonic, ecdsaAccount, err)

	// 测试的错误助记词
	test_mnemonic := "This is a test"
	ecdsaAccount, err = xcc.RetrieveAccountByMnemonic(test_mnemonic, rand.SimplifiedChinese)
	log.Printf("retrieve account from test mnemonic: [%v], ecdsaAccount is %v and err is %v", test_mnemonic, ecdsaAccount, err)

	// NIST助记词
	test_mnemonic = "呈 仓 冯 滚 刚 伙 此 丈 锅 语 揭 弃 精 塘 界 戴 玩 爬 奶 滩 哀 极 样 费"
	ecdsaAccount, err = xcc.RetrieveAccountByMnemonic(test_mnemonic, rand.SimplifiedChinese)
	log.Printf("retrieve account from NIST mnemonic: [%v], ecdsaAccount is %v and err is %v", test_mnemonic, ecdsaAccount, err)
	log.Printf("NIST address is %v and length is %v", ecdsaAccount.Address, len([]rune(ecdsaAccount.Address)))

	// 使用椭圆曲线私钥来签名
	msg := []byte("Welcome to the world of super chain using NIST.")
	strJsonPrivateKey = ecdsaAccount.JsonPrivateKey
	privateKey, err := xcc.GetEcdsaPrivateKeyFromJson([]byte(strJsonPrivateKey))
	// 保留第1个NIST的privateKey
	privateKey1 := privateKey

	sigma, err := xcc.SignECDSA(privateKey, msg)
	log.Printf("sigma is %v and err is %v", sigma, err)

	// 使用椭圆曲线公钥来验证签名是否正确
	//	isSignatureMatch, err = xcc.VerifyECDSA(&privateKey.PublicKey, sigma, msg)
	var ecdsaKeys2 []*ecdsa.PublicKey
	ecdsaKeys2 = append(ecdsaKeys2, &privateKey.PublicKey)
	isSignatureMatch, err := xcc.VerifyXuperSignature(ecdsaKeys2, sigma, msg)
	log.Printf("Verifying & Unmashalling old NIST signature by VerifyXuperSignature, isSignatureMatch is %v and err is %v", isSignatureMatch, err)
	isSignatureMatch, err = xcc.VerifyECDSA(&privateKey.PublicKey, sigma, msg)
	log.Printf("Verifying & Unmashalling old NIST signature by VerifyECDSA, isSignatureMatch is %v and err is %v", isSignatureMatch, err)

	sigma, err = xcc.SignV2ECDSA(privateKey, msg)
	log.Printf("sigma is %s and err is %v", sigma, err)
	isSignatureMatch, err = xcc.VerifyXuperSignature(ecdsaKeys2, sigma, msg)
	log.Printf("Verifying & Unmashalling V2 NIST signature by VerifyXuperSignature, isSignatureMatch is %v and err is %v", isSignatureMatch, err)

	// 创建新的账户，并导出相关文件（含助记词）。生成如下几个文件：1.助记词，2.私钥，3.公钥，4.钱包地址
	err = xcc.ExportNewAccountWithMnemonic("./", rand.SimplifiedChinese, account.StrengthHard)
	if err != nil {
		log.Printf("err happened when ExportNewAccountWithMnemonic: %v", err)
	}

	// 从上一步导出的私钥文件读取私钥
	privateKey, err = xcc.GetEcdsaPrivateKeyFromFile("./private.key")
	// 保留第二个GM的privateKey
	privateKey2 := privateKey

	jsonPrivateKey2, err := json.Marshal(privateKey)
	log.Printf("privateKey has been imported from file, the json format is %s", jsonPrivateKey2)

	// 从上一步导出的公钥文件读取公钥
	publicKey, err := xcc.GetEcdsaPublicKeyFromFile("./public.key")
	jsonPublicKey2, err := json.Marshal(publicKey)
	log.Printf("publicKey has been imported from file, the json format is %s", jsonPublicKey2)

	// 验证钱包地址是否和指定的公钥match。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
	address, _ := xcc.GetAddressFromPublicKey(&privateKey.PublicKey)
	chkResult, nVersion := xcc.VerifyAddressUsingPublicKey(address, &privateKey.PublicKey)
	log.Printf("address check result is %v and the address version is %v", chkResult, nVersion)

	// 验证钱包地址是否是合法的格式。如果成功，返回true和对应的版本号；如果失败，返回false和默认的版本号0
	chkResult, nVersion = xcc.CheckAddressFormat(address)
	log.Printf("address format check result is %v and the address version is %v", chkResult, nVersion)

	// 测试不合法的地址能否被正确的检测出来
	addressFake := "1234567890asdfghjklzxcvbnm"
	chkResult, nVersion = xcc.VerifyAddressUsingPublicKey(addressFake, &privateKey.PublicKey)
	log.Printf("addressFake check result is %v and the address version is %v", chkResult, nVersion)
	chkResult, nVersion = xcc.CheckAddressFormat(addressFake)
	log.Printf("addressFake format check result is %v and the address version is %v", chkResult, nVersion)

	msg = []byte("Hello ecies!")
	ct, err := xcc.EncryptByEcdsaKey(&privateKey.PublicKey, msg)
	if err != nil {
		log.Printf("ECIES encrypt failed and err is: %v", err)
		return
	}

	pt, err := xcc.DecryptByEcdsaKey(privateKey, ct)
	if err != nil {
		log.Printf("ECIES decrypt failed and err is: %v", err)
		return
	}
	log.Printf("pt msg is: %s", pt)

	// 开始算多重签名sig1
	var keys []*ecdsa.PrivateKey
	keys = append(keys, privateKey1)
	keys = append(keys, privateKey2)

	sig1, err := xcc.MultiSign(keys, msg)
	log.Printf("generate multi sig is: %s and err is %v", sig1, err)

	// 开始验证多重签名
	var keys2 []*ecdsa.PublicKey
	keys2 = append(keys2, &privateKey1.PublicKey)
	keys2 = append(keys2, &privateKey2.PublicKey)

	//	chkResult, err = xcc.VerifyMultiSig(keys2, sig1, msg)
	chkResult, err = xcc.VerifyXuperSignature(keys2, sig1, msg)
	log.Printf("VerifyXuperSignature, result is: %v and err is %v", chkResult, err)

	// 验证Schnorr签名算法
	sigma, err = xcc.SignSchnorr(privateKey, msg)
	log.Printf("Schnorr signature is %s and err is %v", sigma, err)

	var keysSchnorr []*ecdsa.PublicKey
	keysSchnorr = append(keysSchnorr, &privateKey.PublicKey)

	//	isSignatureMatch, err = xcc.VerifySchnorr(&privateKey.PublicKey, sigma, msg)
	isSignatureMatch, err = xcc.VerifyXuperSignature(keysSchnorr, sigma, msg)
	log.Printf("Verifying & Unmashalling Schnorr signature, isSignatureMatch is %v and err is %v", isSignatureMatch, err)

	// 验证Schnorr 环签名算法
	log.Printf("Schnorr ring sign ----------")
	log.Printf("keys2 is [%v]:", keys2)
	ringSig, err := xcc.SignSchnorrRing(keys2, originalPrivateKey, msg)
	//	jsonRingSig, _ := json.Marshal(ringSig)
	//	log.Printf("Schnorr ring signature is %s and err is %v", jsonRingSig, err)
	log.Printf("Schnorr ring signature is %s and err is %v", ringSig, err)

	var keysSchnorrRing []*ecdsa.PublicKey
	keysSchnorrRing = append(keysSchnorrRing, &privateKey2.PublicKey)
	keysSchnorrRing = append(keysSchnorrRing, &privateKey1.PublicKey)
	keysSchnorrRing = append(keysSchnorrRing, &originalPrivateKey.PublicKey)
	log.Printf("keysSchnorrRing is [%v]:", keysSchnorrRing)
	//	isSignatureMatch, err = xcc.VerifySchnorrRing(ringSig, msg)
	isSignatureMatch, err = xcc.VerifyXuperSignature(keysSchnorrRing, ringSig, msg)
	log.Printf("Verifying & Unmashalling Schnorr ring signature, isSignatureMatch is %v and err is %v", isSignatureMatch, err)

	// 生成环签名地址
	ringAddress, err := xcc.GetAddressFromPublicKeys(keysSchnorrRing)
	log.Printf("Schnorr ring signature address is %s and err is %v", ringAddress, err)
	isAddressValid, _ := xcc.VerifyAddressUsingPublicKeys(ringAddress, keysSchnorrRing)
	log.Printf("Schnorr ring signature address[%s] is %v", ringAddress, isAddressValid)

	// 生成单个签名地址
	singleAddress, err := xcc.GetAddressFromPublicKey(&privateKey2.PublicKey)
	log.Printf("Schnorr signature address is %s and err is %v", singleAddress, err)
	isAddressValid, _ = xcc.VerifyAddressUsingPublicKey(singleAddress, &privateKey2.PublicKey)
	log.Printf("Simple signature address[%s] is %v", singleAddress, isAddressValid)

	//	return

	// --- hd crypto api ---
	log.Printf("hd crypto api ----------")

	hdMnemonic := "呈 仓 冯 滚 刚 伙 此 丈 锅 语 揭 弃 精 塘 界 戴 玩 爬 奶 滩 哀 极 样 费"
	// 中心化控制中心产生根密钥
	rootKey, _ := xcc.GenerateMasterKeyByMnemonic(hdMnemonic, rand.SimplifiedChinese)
	// 中心化控制中心产生父私钥
	parentPrivateKey, _ := xcc.GenerateChildKey(rootKey, hdapi.HardenedKeyStart+8)
	// 中心化控制中心产生父公钥，并分发给客户端
	parentPublicKey, _ := xcc.ConvertPrvKeyToPubKey(parentPrivateKey)

	hdMsg := "Hello hd msg!"

	// 客户端为每次加密产生子公钥
	newChildPublicKey, err := xcc.GenerateChildKey(parentPublicKey, 18)
	log.Printf("newChildPublicKey is %v and err is %v", newChildPublicKey, err)
	// 客户端使用子公钥加密，产生密文
	//	cryptoMsg, err := hdapi.Encrypt(newChildPublicKey, hdMsg)
	cryptoMsg, err := xcc.EncryptByHdKey(newChildPublicKey, hdMsg)
	log.Printf("cryptoMsg generate err is %v", err)
	log.Printf("cryptoMsg is %v", []byte(cryptoMsg))

	//	cryptoMsg, err = xcc.EncryptByHdKey(newChildPublicKey, hdMsg)
	//	log.Printf("cryptoMsg is %v", []byte(cryptoMsg))

	// 中心化控制中心使用根密钥、子公钥、密文，解密出原文
	realMsg, err := xcc.DecryptByHdKey(newChildPublicKey, rootKey, cryptoMsg)
	log.Printf("realMsg decrypted by root key is: [%s] and err is %v", realMsg, err)

	// 全节点使用一级父私钥、二级子公钥、密文，解密出原文
	realMsg, err = xcc.DecryptByHdKey(newChildPublicKey, parentPrivateKey, cryptoMsg)
	log.Printf("realMsg decrypted by parent private key is: [%s] and err is %v", realMsg, err)

	log.Printf("hd crypto api end----------")

	// -- hd crypto api end ---

	//	msg = []byte("Welcome to the world of secret share.")
	secretMsg := 2147483647
	//	log.Printf("max int is %d", int(^uint32(0)>>1))
	log.Printf("secret_share secret is %d", secretMsg)
	totalShareNumber := 7
	minimumShareNumber := 3

	// ---- ComplexSecret
	log.Printf("----------------------")
	//	complexSecretMsg := []byte("Welcome to the world of secret share.")
	//	complexSecretMsg := []byte("Welcome to the world of secret share 12345678.")
	// 不能太大，否则会由于超出有限域范围产生数据丢失
	complexSecretBigInt, _ := big.NewInt(0).SetString("46950706858566910898749443079945704619073424918969090195410042982588925721159", 0)
	complexSecretMsg := complexSecretBigInt.Bytes()
	//	complexSecretMsg := []byte("a")
	//	log.Printf("secret_share complexSecretMsg is: %s", complexSecretMsg)
	log.Printf("secret_share complexSecretMsg is: %d", complexSecretBigInt)

	complexShares, err := complex_secret_share.ComplexSecretSplit(totalShareNumber, minimumShareNumber, complexSecretMsg)
	log.Printf("secret_share ComplexSecretSplit result is %v and err is %v", complexShares, err)

	retrieveComplexShares := make(map[int]*big.Int, minimumShareNumber)
	number := 0
	for k, v := range complexShares {
		if number >= minimumShareNumber {
			break
		}
		retrieveComplexShares[k] = v
		number++
	}

	secretBytes, _ := complex_secret_share.ComplexSecretRetrieve(retrieveComplexShares)
	//	log.Printf("secret_share ComplexSecretRetrieve result is: %s", secretBytes)
	log.Printf("secret_share ComplexSecretRetrieve result is: %d", big.NewInt(0).SetBytes(secretBytes))
}
