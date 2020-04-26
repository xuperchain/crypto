package gm

import (
	"github.com/xuperchain/crypto/gm/hash"
)

type GmCryptoClient struct {
}

// --- 哈希算法相关 start ---

// 使用SHA256做单次哈希运算
func (xcc *GmCryptoClient) HashUsingSM3(data []byte) []byte {
	hashResult := hash.HashUsingSM3(data)
	return hashResult
}

// --- 哈希算法相关 end ---
