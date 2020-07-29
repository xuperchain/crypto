package hash

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/ripemd160"

	mimcBn256 "github.com/consensys/gnark/crypto/hash/mimc/bn256"
)

func HashUsingSha256(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	out := h.Sum(nil)

	return out
}

// 执行2次SHA256，这是为了防止SHA256算法被攻破。
func DoubleSha256(data []byte) []byte {
	return HashUsingSha256(HashUsingSha256(data))
}

// Ripemd160，这种hash算法可以缩短长度
func HashUsingRipemd160(data []byte) []byte {
	h := ripemd160.New()
	h.Write(data)
	out := h.Sum(nil)

	return out
}

func HashUsingHmac512(data, key []byte) []byte {
	hmac512 := hmac.New(sha512.New, key)
	hmac512.Write(data)
	out := hmac512.Sum(nil)

	return out
}

func HashUsingMiMC(data, seed []byte) []byte {
	mimc := mimcBn256.NewMiMC(string(seed))
	mimc.Write(data)
	out := mimc.Sum(nil)

	return out
}

func HashUsingDefaultMiMC(data []byte) []byte {
	mimc := mimcBn256.NewMiMC("seed")
	mimc.Write(data)
	out := mimc.Sum(nil)

	return out
}
