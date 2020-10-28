package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"log"
	"math/big"

	"github.com/xuperchain/crypto/client/service/xchain"
	"github.com/xuperchain/crypto/core/account"
	//	"github.com/xuperchain/crypto/core/common"
	"github.com/xuperchain/crypto/core/hdwallet/rand"
	//	"github.com/xuperchain/crypto/core/schnorr_sign_new"

	"github.com/xuperchain/crypto/common/math/ecc"
	//	"github.com/xuperchain/crypto/core/threshold/schnorr/dkg"
	//	"github.com/xuperchain/crypto/core/threshold/schnorr/tss_sign"

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
	originalPrivateKey, err := xcc.GetEcdsaPrivateKeyFromJsonStr(strJsonPrivateKey)

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
	privateKey, err := xcc.GetEcdsaPrivateKeyFromJsonStr(strJsonPrivateKey)
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

	// --- 验证new Schnorr签名算法 start ---

	//	sigma, err = schnorr_sign_new.Sign(privateKey, msg)
	//	log.Printf("New Schnorr signature is %s and err is %v", sigma, err)
	//
	//	//	isSignatureMatch, err = xcc.VerifySchnorr(&privateKey.PublicKey, sigma, msg)
	//	isSignatureMatch, err = schnorr_sign_new.Verify(&privateKey.PublicKey, sigma, msg)
	//	log.Printf("Verifying & Unmashalling new Schnorr signature, isSignatureMatch is %v and err is %v", isSignatureMatch, err)

	// --- 验证new Schnorr签名算法 end ---

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
	//	secretMsg := 2147483647
	//	log.Printf("max int is %d", int(^uint32(0)>>1))
	//	log.Printf("secret_share secret is %d", secretMsg)
	totalShareNumber := 7
	minimumShareNumber := 3

	// ---- ComplexSecret ---
	log.Printf("----------------------")
	//	complexSecretMsg := []byte("Welcome to the world of secret share.")
	//	complexSecretMsg := []byte("Welcome to the world of secret share 12345678.")
	// 不能太大，否则会由于超出有限域范围产生数据丢失
	complexSecretBigInt, _ := big.NewInt(0).SetString("46950706858566910898749443079945704619073424918969090195410042982588925721159", 0)
	complexSecretMsg := complexSecretBigInt.Bytes()
	//	complexSecretMsg := []byte("a")
	//	log.Printf("secret_share complexSecretMsg is: %s", complexSecretMsg)
	log.Printf("secret_share complexSecretMsg is: %d", complexSecretBigInt)

	complexShares, err := xcc.SecretSplit(totalShareNumber, minimumShareNumber, complexSecretMsg)
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

	secretBytes, _ := xcc.SecretRetrieve(retrieveComplexShares)
	//	log.Printf("secret_share ComplexSecretRetrieve result is: %s", secretBytes)
	log.Printf("secret_share ComplexSecretRetrieve result is: %d", big.NewInt(0).SetBytes(secretBytes))

	//--- private key split start ---
	strPrivKeyShares, err := xcc.SplitPrivateKey(strJsonPrivateKey, totalShareNumber, minimumShareNumber)
	log.Printf("share_key SplitPrivateKey result is: %s, and err is: %v", strPrivKeyShares, err)

	jsonPrivKey, err := xcc.RetrievePrivateKeyByShares(strPrivKeyShares[0:minimumShareNumber])
	log.Printf("share_key RetrievedPrivateKey fragments are: %s, result is: %s, and err is: %v", strPrivKeyShares, jsonPrivKey, err)
	//--- private key split end ---

	// --- zkp start ---
	zkpInfo := xcc.ZkpSetupMiMC()
	secretMsg := []byte("test for zkp")
	proof, err := xcc.ZkpProveMiMC(zkpInfo.R1CS, zkpInfo.ProvingKey, secretMsg)
	if err != nil {
		log.Printf("ZkpProveMiMC failed and err is: %v", err)
		return
	}
	log.Printf("ZkpProveMiMC proof is: %v", proof)

	hashResult = xcc.HashUsingDefaultMiMC(secretMsg)
	verifyResult, err := xcc.ZkpVerifyMiMC(proof, zkpInfo.VerifyingKey, hashResult)
	log.Printf("verifyResult proof is: %v and err is: %v", verifyResult, err)

	verifyResult, err = xcc.ZkpVerifyMiMC(proof, zkpInfo.VerifyingKey, []byte("test2 for zkp"))
	log.Printf("verifyResult proof is: %v and err is: %v", verifyResult, err)

	// --- zkp end ---

	// --- 验证门限签名 start ---

	log.Printf("threhold sig start...")

	// 开始为多方建立初始化信息
	//	partnerPublic1 := &dkg.PartnerPublic{Index: 1, IndentityKey: big.NewInt(0).SetBytes([]byte("threhold partner 1"))}
	//	partnerPublic2 := &dkg.PartnerPublic{Index: 2, IndentityKey: big.NewInt(0).SetBytes([]byte("threhold partner 2"))}
	//	partnerPublic3 := &dkg.PartnerPublic{Index: 3, IndentityKey: big.NewInt(0).SetBytes([]byte("threhold partner 3"))}

	//	log.Printf("partnerPublic1 is %v", partnerPublic1)

	//	shares1, verifyPoints1, err := dkg.SecretShareLocalKeyGenerateWithVerifyPoints(3, 2, big.NewInt(1).Bytes())
	//	shares2, verifyPoints2, err := dkg.SecretShareLocalKeyGenerateWithVerifyPoints(3, 2, big.NewInt(2).Bytes())
	//	shares3, verifyPoints3, err := dkg.SecretShareLocalKeyGenerateWithVerifyPoints(3, 2, big.NewInt(3).Bytes())

	//	log.Printf("shares1 is %v, verifyPoints1 is %v and err is %v", shares1, verifyPoints1, err)
	//	log.Printf("shares2 is %v, verifyPoints2 is %v and err is %v", shares2, verifyPoints2, err)
	//	log.Printf("shares3 is %v, verifyPoints3 is %v and err is %v", shares3, verifyPoints3, err)

	//	partnerShares1 := &dkg.PartnerShares{PartnerInfo: partnerPublic1, Shares: shares1, VerifyPoints: verifyPoints1}
	//	partnerShares2 := &dkg.PartnerShares{PartnerInfo: partnerPublic2, Shares: shares2, VerifyPoints: verifyPoints2}
	//	partnerShares3 := &dkg.PartnerShares{PartnerInfo: partnerPublic3, Shares: shares3, VerifyPoints: verifyPoints3}
	//
	//	log.Printf("partnerShares1 is %v", partnerShares1)

	//	//	allPartnerShares := make([]*dkg.PartnerShares, 3)
	//	var allPartnerShares []*dkg.PartnerShares
	//	allPartnerShares = append(allPartnerShares, partnerShares1)
	//	allPartnerShares = append(allPartnerShares, partnerShares2)
	//	allPartnerShares = append(allPartnerShares, partnerShares3)
	//
	//	// 从收集的所有碎片中保存与自己相关的密钥部分，并保存在本地
	//	localShares1 := dkg.SecretShareLocalKeyGather(allPartnerShares, 1)
	//	localShares2 := dkg.SecretShareLocalKeyGather(allPartnerShares, 2)
	//	localShares3 := dkg.SecretShareLocalKeyGather(allPartnerShares, 3)

	//	// 从本地share计算自己的秘密Xi
	//	localXi1 := dkg.CalcuateXi(localShares1)
	//	localXi2 := dkg.CalcuateXi(localShares2)
	//	localXi3 := dkg.CalcuateXi(localShares3)
	//
	//	log.Printf("localXi[N] is: %v, %v, %v", localXi1, localXi2, localXi3)

	//	// 计算公共验证点部分，并保存在本地
	//	verifyPoints, err := dkg.SecretShareVerifyPointsGather(allPartnerShares, 2)
	//	jsonVerifyPoints, _ := json.Marshal(verifyPoints)
	//	log.Printf("verifyPoints is: %s and err is %v", jsonVerifyPoints, err)
	//
	//	// 计算公共public key point
	//	publicKeyPoint, err := dkg.CalculatePublicKey(verifyPoints)
	//	jsonPublicKeyPoint, _ := json.Marshal(publicKeyPoint)
	//	log.Printf("publicKeyPoint is: %s and err is %v", jsonPublicKeyPoint, err)

	//	// 计算每一方的public key point
	//	publicKeyPoints, err := dkg.CalculatePublicKeys(verifyPoints, allPartnerShares, 2)
	//	jsonPublicKeyPoints, _ := json.Marshal(publicKeyPoints)
	//	log.Printf("publicKeyPoints is: %s and err is %v", jsonPublicKeyPoints, err)

	// -----
	//	localPrivateKey1 := new(ecdsa.PrivateKey)
	//	localPrivateKey1.PublicKey.Curve = publicKeyPoint.Curve //elliptic.P256()
	//	localPrivateKey1.D = localXi1
	//	localPrivateKey1.PublicKey.X, localPrivateKey1.PublicKey.Y = publicKeyPoint.Curve.ScalarBaseMult(localXi1.Bytes())
	//
	//	localPrivateKey2 := new(ecdsa.PrivateKey)
	//	localPrivateKey2.PublicKey.Curve = publicKeyPoint.Curve //elliptic.P256()
	//	localPrivateKey2.D = localXi2
	//	localPrivateKey2.PublicKey.X, localPrivateKey1.PublicKey.Y = publicKeyPoint.Curve.ScalarBaseMult(localXi2.Bytes())
	//
	//	localPrivateKey3 := new(ecdsa.PrivateKey)
	//	localPrivateKey3.PublicKey.Curve = publicKeyPoint.Curve //elliptic.P256()
	//	localPrivateKey3.D = localXi3
	//	localPrivateKey3.PublicKey.X, localPrivateKey1.PublicKey.Y = publicKeyPoint.Curve.ScalarBaseMult(localXi3.Bytes())
	//
	//	jsonLocalPrivateKey3, _ := json.Marshal(localPrivateKey3)
	//	log.Printf("localPrivateKey3 is: %s", jsonLocalPrivateKey3, err)

	// ---
	// --- DKG ---
	// 每一方生成自己的秘密碎片
	// 3个潜在参与节点，门限要求是2，也就是要大于等于2个节点参与才能形成有效签名
	shares1, verifyPoints1, _ := xcc.GetLocalShares(3, 2)
	shares2, verifyPoints2, _ := xcc.GetLocalShares(3, 2)
	shares3, verifyPoints3, _ := xcc.GetLocalShares(3, 2)

	// 碎片交换
	var localShares1 []*big.Int
	var localShares2 []*big.Int
	var localShares3 []*big.Int

	// 节点编号1
	localShares1 = append(localShares1, shares1[1])
	localShares1 = append(localShares1, shares2[1])
	localShares1 = append(localShares1, shares3[1])
	log.Printf("localShares1 is: %v", localShares1)

	// 节点编号2
	localShares2 = append(localShares2, shares1[2])
	localShares2 = append(localShares2, shares2[2])
	localShares2 = append(localShares2, shares3[2])
	log.Printf("localShares2 is: %v", localShares2)

	// 节点编号3
	localShares3 = append(localShares3, shares1[3])
	localShares3 = append(localShares3, shares2[3])
	localShares3 = append(localShares3, shares3[3])
	log.Printf("localShares3 is: %v", localShares3)

	// 计算本地私钥
	localPrivateKey1 := xcc.GetLocalPrivateKeyByShares(localShares1)
	localPrivateKey2 := xcc.GetLocalPrivateKeyByShares(localShares2)
	localPrivateKey3 := xcc.GetLocalPrivateKeyByShares(localShares3)

	jsonLocalPrivateKey3, _ := json.Marshal(localPrivateKey3)
	log.Printf("localPrivateKey3 is: %s", jsonLocalPrivateKey3)

	// 验证点交换
	var verifyPoints []*ecc.Point

	verifyPoints = append(verifyPoints, verifyPoints1[0])
	verifyPoints = append(verifyPoints, verifyPoints2[0])
	verifyPoints = append(verifyPoints, verifyPoints3[0])

	// 计算公钥
	tssPublickey, _ := xcc.GetSharedPublicKey(verifyPoints)
	jsonTssPublickey, _ := json.Marshal(tssPublickey)
	log.Printf("tssPublickey is: %s", jsonTssPublickey)

	// --- DSG & 验证 ---

	//	var tssKeys []*ecdsa.PrivateKey
	//	tssKeys = append(tssKeys, localPrivateKey1)
	//	tssKeys = append(tssKeys, localPrivateKey2)

	// tss签名
	rk1, _ := xcc.GetRandom32Bytes()
	rk2, _ := xcc.GetRandom32Bytes()
	r1 := xcc.GetRiUsingRandomBytes(tssPublickey, rk1)
	r2 := xcc.GetRiUsingRandomBytes(tssPublickey, rk2)

	var arrayOfRi [][]byte
	arrayOfRi = append(arrayOfRi, r1)
	arrayOfRi = append(arrayOfRi, r2)

	r := xcc.GetRUsingAllRi(tssPublickey, arrayOfRi)

	var ks []*big.Int
	ks = append(ks, big.NewInt(1)) // 节点编号1
	ks = append(ks, big.NewInt(2)) // 节点编号2
	// 本次节点编号3不参与签名过程
	//	ks = append(ks, big.NewInt(3)) // 节点编号3

	w1 := xcc.GetXiWithcoef(ks, 0, localPrivateKey1)
	w2 := xcc.GetXiWithcoef(ks, 1, localPrivateKey2)

	c := elliptic.Marshal(tssPublickey.Curve, tssPublickey.X, tssPublickey.Y)

	s1 := xcc.GetSiUsingKCRMWithCoef(rk1, c, r, msg, w1)
	s2 := xcc.GetSiUsingKCRMWithCoef(rk2, c, r, msg, w2)

	var arrayOfSi [][]byte
	arrayOfSi = append(arrayOfSi, s1)
	arrayOfSi = append(arrayOfSi, s2)

	s := xcc.GetSUsingAllSi(arrayOfSi)
	//	log.Printf("all of s is: %d", big.NewInt(0).SetBytes(s))

	tssSig, _ := xcc.GenerateTssSignSignature(s, r)
	log.Printf("tssSig is: %s", tssSig)

	// 验证tss签名
	var tssPublicKeys []*ecdsa.PublicKey
	tssPublicKeys = append(tssPublicKeys, tssPublickey)

	chkResult, _ = xcc.VerifyXuperSignature(tssPublicKeys, tssSig, msg)
	log.Printf("verify tss sig chkResult is: %v", chkResult)

	log.Printf("threhold sig end...")

	// --- 验证门限签名 end ---

	// --- 验证分层门限签名 start ---

	log.Printf("Hierarchical threhold sig start...")

	// 目标：团队中3个员工和3个经理，要求，3个团队成员参与才能获得授权签名，其中至少要有2个经理配合
	// 低授权级别的门限要求必须要大于高授权级别的门限要求，例如，纯成员级别以上是3/6，那么纯经理级别以上只能是2/3，因为2要小于3。
	// 3个经理的节点编号：1、2、3
	// 3个员工的节点编号：4、5、6

	// --- DKG ---
	// 员工（级别E部分）和经理（级别M部分）每一方生成自己的秘密碎片
	// 6个潜在参与节点，门限要求是3，也就是要大于等于3个节点参与才能形成有效签名
	sharesE1, verifyPointsE1, _ := xcc.GetLocalShares(6, 3)
	sharesE2, verifyPointsE2, _ := xcc.GetLocalShares(6, 3)
	sharesE3, verifyPointsE3, _ := xcc.GetLocalShares(6, 3)
	sharesE4, verifyPointsE4, _ := xcc.GetLocalShares(6, 3)
	sharesE5, verifyPointsE5, _ := xcc.GetLocalShares(6, 3)
	sharesE6, verifyPointsE6, _ := xcc.GetLocalShares(6, 3)

	// 碎片交换
	var localSharesE1 []*big.Int
	var localSharesE2 []*big.Int
	var localSharesE3 []*big.Int
	var localSharesE4 []*big.Int
	var localSharesE5 []*big.Int
	var localSharesE6 []*big.Int

	// 员工节点编号1的E部分
	localSharesE1 = append(localSharesE1, sharesE1[1])
	localSharesE1 = append(localSharesE1, sharesE2[1])
	localSharesE1 = append(localSharesE1, sharesE3[1])
	localSharesE1 = append(localSharesE1, sharesE4[1])
	localSharesE1 = append(localSharesE1, sharesE5[1])
	localSharesE1 = append(localSharesE1, sharesE6[1])
	log.Printf("localSharesE1 is: %v", localSharesE1)

	// 员工节点编号2的E部分
	localSharesE2 = append(localSharesE2, sharesE1[2])
	localSharesE2 = append(localSharesE2, sharesE2[2])
	localSharesE2 = append(localSharesE2, sharesE3[2])
	localSharesE2 = append(localSharesE2, sharesE4[2])
	localSharesE2 = append(localSharesE2, sharesE5[2])
	localSharesE2 = append(localSharesE2, sharesE6[2])
	log.Printf("localSharesE2 is: %v", localSharesE2)

	// 员工节点编号3的E部分
	localSharesE3 = append(localSharesE3, sharesE1[3])
	localSharesE3 = append(localSharesE3, sharesE2[3])
	localSharesE3 = append(localSharesE3, sharesE3[3])
	localSharesE3 = append(localSharesE3, sharesE4[3])
	localSharesE3 = append(localSharesE3, sharesE5[3])
	localSharesE3 = append(localSharesE3, sharesE6[3])
	log.Printf("localSharesE3 is: %v", localSharesE3)

	// 员工节点编号4的E部分
	localSharesE4 = append(localSharesE4, sharesE1[4])
	localSharesE4 = append(localSharesE4, sharesE2[4])
	localSharesE4 = append(localSharesE4, sharesE3[4])
	localSharesE4 = append(localSharesE4, sharesE4[4])
	localSharesE4 = append(localSharesE4, sharesE5[4])
	localSharesE4 = append(localSharesE4, sharesE6[4])
	log.Printf("localSharesE4 is: %v", localSharesE4)

	// 员工节点编号5的E部分
	localSharesE5 = append(localSharesE5, sharesE1[5])
	localSharesE5 = append(localSharesE5, sharesE2[5])
	localSharesE5 = append(localSharesE5, sharesE3[5])
	localSharesE5 = append(localSharesE5, sharesE4[5])
	localSharesE5 = append(localSharesE5, sharesE5[5])
	localSharesE5 = append(localSharesE5, sharesE6[5])
	log.Printf("localSharesE5 is: %v", localSharesE5)

	// 员工节点编号6的E部分
	localSharesE6 = append(localSharesE6, sharesE1[6])
	localSharesE6 = append(localSharesE6, sharesE2[6])
	localSharesE6 = append(localSharesE6, sharesE3[6])
	localSharesE6 = append(localSharesE6, sharesE4[6])
	localSharesE6 = append(localSharesE6, sharesE5[6])
	localSharesE6 = append(localSharesE6, sharesE6[6])
	log.Printf("localSharesE6 is: %v", localSharesE6)

	//--经理（级别M部分）每一方生成自己的秘密碎片
	// 注意，经理同时持有M部分和E部分的碎片
	// 3个潜在参与节点，门限要求是2，也就是要大于等于2个节点参与才能形成有效签名
	// 员工节点编号1的M部分
	sharesM1, verifyPointsM1, _ := xcc.GetLocalShares(3, 2)
	sharesM2, verifyPointsM2, _ := xcc.GetLocalShares(3, 2)
	sharesM3, verifyPointsM3, _ := xcc.GetLocalShares(3, 2)

	// 碎片交换
	var localSharesM1 []*big.Int
	var localSharesM2 []*big.Int
	var localSharesM3 []*big.Int

	// 节点编号1的M部分
	localSharesM1 = append(localSharesM1, sharesM1[1])
	localSharesM1 = append(localSharesM1, sharesM2[1])
	localSharesM1 = append(localSharesM1, sharesM3[1])
	log.Printf("localSharesM1 is: %v", localSharesM1)

	// 节点编号2的E部分
	localSharesM2 = append(localSharesM2, sharesM1[2])
	localSharesM2 = append(localSharesM2, sharesM2[2])
	localSharesM2 = append(localSharesM2, sharesM3[2])
	log.Printf("localSharesM2 is: %v", localSharesM2)

	// 节点编号3的E部分
	localSharesM3 = append(localSharesM3, sharesM1[3])
	localSharesM3 = append(localSharesM3, sharesM2[3])
	localSharesM3 = append(localSharesM3, sharesM3[3])
	log.Printf("localSharesM3 is: %v", localSharesM3)

	// 验证点交换
	var verifyPointsEM []*ecc.Point

	verifyPointsEM = append(verifyPointsEM, verifyPointsE1[0])
	verifyPointsEM = append(verifyPointsEM, verifyPointsE2[0])
	verifyPointsEM = append(verifyPointsEM, verifyPointsE3[0])
	verifyPointsEM = append(verifyPointsEM, verifyPointsE4[0])
	verifyPointsEM = append(verifyPointsEM, verifyPointsE5[0])
	verifyPointsEM = append(verifyPointsEM, verifyPointsE6[0])

	verifyPointsEM = append(verifyPointsEM, verifyPointsM1[0])
	verifyPointsEM = append(verifyPointsEM, verifyPointsM2[0])
	verifyPointsEM = append(verifyPointsEM, verifyPointsM3[0])

	// 计算公钥
	tssPublickeyEM, _ := xcc.GetSharedPublicKey(verifyPointsEM)
	jsonTssPublickeyEM, _ := json.Marshal(tssPublickeyEM)
	log.Printf("tssPublickeyEM is: %s", jsonTssPublickeyEM)

	c = elliptic.Marshal(tssPublickeyEM.Curve, tssPublickeyEM.X, tssPublickeyEM.Y)

	// 计算本地私钥的E部分
	localPrivateKeyE1 := xcc.GetLocalPrivateKeyByShares(localSharesE1)
	localPrivateKeyE2 := xcc.GetLocalPrivateKeyByShares(localSharesE2)
	localPrivateKeyE3 := xcc.GetLocalPrivateKeyByShares(localSharesE3)
	localPrivateKeyE4 := xcc.GetLocalPrivateKeyByShares(localSharesE4)
	localPrivateKeyE5 := xcc.GetLocalPrivateKeyByShares(localSharesE5)
	localPrivateKeyE6 := xcc.GetLocalPrivateKeyByShares(localSharesE6)

	jsonLocalPrivateKeyE3, _ := json.Marshal(localPrivateKeyE3)
	log.Printf("localPrivateKeyE3 is: %s", jsonLocalPrivateKeyE3)
	jsonLocalPrivateKeyE5, _ := json.Marshal(localPrivateKeyE5)
	log.Printf("localPrivateKeyE5 is: %s", jsonLocalPrivateKeyE5)
	jsonLocalPrivateKeyE6, _ := json.Marshal(localPrivateKeyE6)
	log.Printf("localPrivateKeyE6 is: %s", jsonLocalPrivateKeyE6)

	// 计算本地私钥的M部分
	localPrivateKeyM1 := xcc.GetLocalPrivateKeyByShares(localSharesM1)
	localPrivateKeyM2 := xcc.GetLocalPrivateKeyByShares(localSharesM2)
	localPrivateKeyM3 := xcc.GetLocalPrivateKeyByShares(localSharesM3)

	jsonLocalPrivateKeyM3, _ := json.Marshal(localPrivateKeyM3)
	log.Printf("localPrivateKeyM3 is: %s", jsonLocalPrivateKeyM3)

	// --- DSG & 验证 ---

	//	var tssKeys []*ecdsa.PrivateKey
	//	tssKeys = append(tssKeys, localPrivateKey1)
	//	tssKeys = append(tssKeys, localPrivateKey2)

	// tss签名
	rk1, _ = xcc.GetRandom32Bytes()
	rk2, _ = xcc.GetRandom32Bytes()
	rk4, _ := xcc.GetRandom32Bytes()

	r1 = xcc.GetRiUsingRandomBytes(tssPublickeyEM, rk1)
	r2 = xcc.GetRiUsingRandomBytes(tssPublickeyEM, rk2)
	r4 := xcc.GetRiUsingRandomBytes(tssPublickeyEM, rk4)

	//	var arrayOfRiM [][]byte
	//	arrayOfRiM = append(arrayOfRiM, r1)
	//	arrayOfRiM = append(arrayOfRiM, r2)
	//
	//	rM := xcc.GetRUsingAllRi(tssPublickey, arrayOfRiM)
	//
	//	var arrayOfRiE [][]byte
	//	arrayOfRiE = append(arrayOfRiE, r1)
	//	arrayOfRiE = append(arrayOfRiE, r2)
	//	arrayOfRiE = append(arrayOfRiE, r4)
	//
	//	rE := xcc.GetRUsingAllRi(tssPublickeyEM, arrayOfRiE)

	// 通用R
	var arrayOfRiEM [][]byte
	arrayOfRiEM = append(arrayOfRiEM, r1)
	arrayOfRiEM = append(arrayOfRiEM, r2)
	arrayOfRiEM = append(arrayOfRiEM, r4)

	r = xcc.GetRUsingAllRi(tssPublickeyEM, arrayOfRiEM)

	// SiM部分计算
	var ksEM []*big.Int
	ksEM = append(ksEM, big.NewInt(1)) // 节点编号1
	ksEM = append(ksEM, big.NewInt(2)) // 节点编号2

	wM1 := xcc.GetXiWithcoef(ksEM, 0, localPrivateKeyM1)
	wM2 := xcc.GetXiWithcoef(ksEM, 1, localPrivateKeyM2)

	//	sM1 := xcc.GetSiUsingKCRMWithCoef(rk1, c, r, msg, wM1)
	//	sM2 := xcc.GetSiUsingKCRMWithCoef(rk2, c, r, msg, wM2)
	sM1 := xcc.GetSiUsingKCRMWithCoefNoKi(c, r, msg, wM1)
	sM2 := xcc.GetSiUsingKCRMWithCoefNoKi(c, r, msg, wM2)

	// SiE部分计算
	ksEM = append(ksEM, big.NewInt(4)) // 节点编号4

	wE1 := xcc.GetXiWithcoef(ksEM, 0, localPrivateKeyE1)
	wE2 := xcc.GetXiWithcoef(ksEM, 1, localPrivateKeyE2)
	wE4 := xcc.GetXiWithcoef(ksEM, 2, localPrivateKeyE4) // 对应节点编号4

	sE1 := xcc.GetSiUsingKCRMWithCoef(rk1, c, r, msg, wE1)
	sE2 := xcc.GetSiUsingKCRMWithCoef(rk2, c, r, msg, wE2)
	sE4 := xcc.GetSiUsingKCRMWithCoef(rk4, c, r, msg, wE4)

	// 计算arrayOfSi(每个节点的M部分和E部分之和)
	var arrayOfSiEM [][]byte
	arrayOfSiEM = append(arrayOfSiEM, sM1)
	arrayOfSiEM = append(arrayOfSiEM, sM2)
	arrayOfSiEM = append(arrayOfSiEM, sE1)
	arrayOfSiEM = append(arrayOfSiEM, sE2)
	arrayOfSiEM = append(arrayOfSiEM, sE4)

	s = xcc.GetSUsingAllSi(arrayOfSiEM)
	//	log.Printf("all of s is: %d", big.NewInt(0).SetBytes(s))

	tssSig, _ = xcc.GenerateTssSignSignature(s, r)
	log.Printf("tssSig is: %s", tssSig)

	// 验证tss签名
	var tssPublicKeysEM []*ecdsa.PublicKey
	tssPublicKeysEM = append(tssPublicKeysEM, tssPublickeyEM)

	chkResult, _ = xcc.VerifyXuperSignature(tssPublicKeysEM, tssSig, msg)
	log.Printf("verify tss sig chkResult is: %v", chkResult)

	log.Printf("Hierarchical threhold sig end...")

	// --- 验证分层门限签名 end ---

	// --- 验证BLS签名 start ---

	// 生成BLS密钥对
	privateKeyBLS, publicKeyBLS := xcc.GenerateBlsKeyPair()

	// 验证BLS签名算法
	blsSig, err := xcc.SignBls(privateKeyBLS, msg)
	log.Printf("BLS signature is %s and err is %v", blsSig, err)

	isSignatureMatch, err = xcc.VerifyBlsSig(publicKeyBLS, blsSig, msg)
	log.Printf("Verifying & Unmashalling BLS signature, isSignatureMatch is %v and err is %v", isSignatureMatch, err)

	_, publicKeyBLS2 := xcc.GenerateBlsKeyPair()
	isSignatureMatch, err = xcc.VerifyBlsSig(publicKeyBLS2, blsSig, msg)
	log.Printf("Verifying & Unmashalling BLS signature using publicKeyBLS2, isSignatureMatch is %v and err is %v", isSignatureMatch, err)

	// --- 验证BLS签名 end ---

	// --- 测试密钥椭圆曲线更换 start ---

	s256PrivateKeyE1 := xcc.ChangePrivCurveToS256k1(localPrivateKeyE1)
	jsonS256PrivateKeyE1, _ := json.Marshal(s256PrivateKeyE1)
	log.Printf("s256PrivateKeyE1 is: %s", jsonS256PrivateKeyE1)

	// --- 测试密钥椭圆曲线更换 end ---
}
