package main

import (
	"log"

	"github.com/xuperchain/crypto/client/service/gm"
)

func main() {
	gcc := new(gm.GmCryptoClient)

	hashResult := gcc.HashUsingSM3([]byte("This is xchain crypto"))
	log.Printf("Hash result for [This is xchain crypto] is: %s", hashResult)
}
