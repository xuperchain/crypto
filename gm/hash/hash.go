package hash

import (
	"crypto/hmac"
	"crypto/sha512"

	"github.com/xuperchain/crypto/gm/gmsm/sm3"

	"golang.org/x/crypto/ripemd160"
)

func HashUsingSM3(data []byte) []byte {
	var h sm3.SM3

	h.Reset()
	h.Write(data)
	out := h.Sum(nil)

	return out
}

// Ripemd160，这种hash算法可以缩短长度
func HashUsingRipemd160(data []byte) []byte {
	h := ripemd160.New()
	h.Write(data)
	out := h.Sum(nil)

	return out
}

func HashUsingHmac512(seed, key []byte) []byte {
	hmac512 := hmac.New(sha512.New, key)
	hmac512.Write(seed)
	out := hmac512.Sum(nil)

	return out
}
