package common

import (
	"math/big"
)

// 定义签名中所包含的标记符的值，及其所对应的签名算法的类型
const (
	// ECDSA签名算法
	ECDSA = "ECDSA"
	// Schnorr签名算法，EDDSA的前身
	Schnorr = "Schnorr"
	// Schnorr环签名算法
	SchnorrRing = "SchnorrRing"
	// 多重签名算法
	MultiSig = "MultiSig"
	// 门限签名算法
	TssSig = "TssSig"
	// Bls签名算法
	BlsSig = "BlsSig"
)

// --- 签名数据结构相关 start ---

// 统一的签名结构
type XuperSignature struct {
	SigType    string
	SigContent []byte
}

// ECDSA签名
type ECDSASignature struct {
	R, S *big.Int
}

// Schnorr签名，EDDSA的前身
type SchnorrSignature struct {
	E, S *big.Int
}

// --- Schnorr环签名的数据结构定义 start ---

type PublicKeyFactor struct {
	X, Y *big.Int
}

// Schnorr环签名
type RingSignature struct {
	CurveName string
	Members   []*PublicKeyFactor
	E         *big.Int
	S         []*big.Int
}

// --- Schnorr环签名的数据结构定义 end ---

// 多重签名
type MultiSignature struct {
	S []byte
	R []byte
}

// 门限签名
type TssSignature struct {
	S []byte
	R []byte
}

// BLS签名
type BlsSignature struct {
	S []byte
}

// --- 签名数据结构相关 end ---
