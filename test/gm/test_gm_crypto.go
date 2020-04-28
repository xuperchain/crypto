package main

import (
	"log"

	"github.com/xuperchain/crypto/client/service/gm"
	"github.com/xuperchain/crypto/client/service/xchain"
	"github.com/xuperchain/crypto/gm/account"
	"github.com/xuperchain/crypto/gm/hdwallet/rand"
)

func main() {
	gcc := new(gm.GmCryptoClient)
	xcc := new(xchain.XchainCryptoClient)

	// --- 哈希算法相关 start ---
	hashResult := gcc.HashUsingSM3([]byte("This is xchain crypto"))
	log.Printf("Hash result for [This is xchain crypto] is: %s", hashResult)
	// --- 哈希算法相关 end ---

	// --- 地址生成相关 start ---
	ecdsaAccount, _ := xcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
	log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)

	// --- 地址生成相关 end ---
}
