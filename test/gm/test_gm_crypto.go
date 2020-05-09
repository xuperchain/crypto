package main

import (
	"log"

	"github.com/xuperchain/crypto/client/service/gm"
	"github.com/xuperchain/crypto/gm/account"
	"github.com/xuperchain/crypto/gm/hdwallet/rand"
)

func main() {
	gcc := new(gm.GmCryptoClient)

	// --- 哈希算法相关 start ---
	hashResult := gcc.HashUsingSM3([]byte("This is xchain crypto"))
	log.Printf("Hash result for [This is xchain crypto] is: %s", hashResult)
	// --- 哈希算法相关 end ---

	// --- 账户生成相关 start ---
	ecdsaAccount, err := gcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
	if err != nil {
		log.Printf("CreateNewAccountWithMnemonic failed and err is: %v", err)
		return
	}
	log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)
	// --- 账户生成相关 end ---

	// --- 账户恢复相关 start ---
	// 从助记词恢复账户
	// 测试错误助记词
	test_mnemonic := "This is a test"
	wrongEcdsaAccount, err := gcc.RetrieveAccountByMnemonic(test_mnemonic, rand.SimplifiedChinese)
	log.Printf("retrieve account from test mnemonic: [%v], ecdsaAccount is %v and err is %v", test_mnemonic, wrongEcdsaAccount, err)

	// 测试正确助记词
	ecdsaAccount, err = gcc.RetrieveAccountByMnemonic(ecdsaAccount.Mnemonic, rand.SimplifiedChinese)
	if err != nil {
		log.Printf("RetrieveAccountByMnemonic failed and err is: %v", err)
		return
	}
	log.Printf("retrieve account from mnemonic %v, ecdsaAccount is %v and err is %v", ecdsaAccount.Mnemonic, ecdsaAccount, err)
	// --- 账户恢复相关 end ---

	// --- ECDSA签名算法相关 start ---
	msg := []byte("Welcome to the world of super chain using GM.")
	strJsonPrivateKey := ecdsaAccount.JsonPrivateKey
	privateKey, err := gcc.GetEcdsaPrivateKeyFromJsonStr(strJsonPrivateKey)
	sig, err := gcc.SignECDSA(privateKey, msg)
	log.Printf("sig is %v and err is %v", sig, err)

	isSignatureMatch, err := gcc.VerifyECDSA(&privateKey.PublicKey, sig, msg)
	log.Printf("Verifying & Unmashalling GM ecdsa signature by VerifyECDSA, isSignatureMatch is %v and err is %v", isSignatureMatch, err)
	// --- ECDSA签名算法相关 end ---

	// --- 非对称加密算法相关 start ---
	msg = []byte("Hello encryption!")
	ct, err := gcc.EncryptByEcdsaKey(&privateKey.PublicKey, msg)
	if err != nil {
		log.Printf("Encrypt failed and err is: %v", err)
		return
	}

	pt, err := gcc.DecryptByEcdsaKey(privateKey, ct)
	if err != nil {
		log.Printf("Decrypt failed and err is: %v", err)
		return
	}
	log.Printf("pt msg after decryption is: %s", pt)
	// --- 非对称加密算法相关 end ---
}
