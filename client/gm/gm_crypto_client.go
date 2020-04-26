package main

import (
	"github.com/xuperchain/crypto/client/service/gm"
)

// GetInstance return the gm client
func GetInstance() interface{} {
	return &gm.GmCryptoClient{}
}
