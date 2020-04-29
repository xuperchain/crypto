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
	ecdsaAccount, _ := gcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
	log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)
	// --- 账户生成相关 end ---

	// --- 账户恢复相关 start ---
	// 从助记词恢复账户
	ecdsaAccount, err := gcc.RetrieveAccountByMnemonic(ecdsaAccount.Mnemonic, rand.SimplifiedChinese)
	if err != nil {
		log.Printf("RetrieveAccountByMnemonic failed and err is: %v", err)
		return
	}
	log.Printf("retrieve account from mnemonic %v, ecdsaAccount is %v and err is %v", ecdsaAccount.Mnemonic, ecdsaAccount, err)

	// 测试的错误助记词
	test_mnemonic := "This is a test"
	ecdsaAccount, err = gcc.RetrieveAccountByMnemonic(test_mnemonic, rand.SimplifiedChinese)
	log.Printf("retrieve account from test mnemonic: [%v], ecdsaAccount is %v and err is %v", test_mnemonic, ecdsaAccount, err)
	// --- 账户恢复相关 end ---
}
