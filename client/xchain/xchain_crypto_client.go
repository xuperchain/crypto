package main

import (
	"github.com/xuperchain/crypto/client/service/xchain"
)

// GetInstance return the default xchain client
func GetInstance() interface{} {
	return &xchain.XchainCryptoClient{}
}
