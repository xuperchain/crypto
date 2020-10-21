package bls_sign

/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com> 西二旗第一帅
*/

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"github.com/xuperchain/crypto/core/common"
	"github.com/xuperchain/crypto/core/hash"

	"github.com/cloudflare/bn256"
)

type PrivateKey struct {
	X *big.Int
}

type PublicKey struct {
	P *bn256.G2
}

// BLS signature uses a particular function, defined as:
// S = pk * H(m)
//
// H is a hash function, for instance SHA256 or SM3.
// S is the signature.
// m is the message to sign.
// pk is the private key, which can be considered as a secret big number.
//
// To verify the signature, check that whether the result of e(P, H(m)) is equal to e(G, S) or not.
// Which means that: e(P, H(m)) = e(G, S)
// G is the base point or the generator point.
// P is the public key = pk*G.
// e is a special elliptic curve pairing function which has this feature: e(x*P, Q) = e(P, x*Q).
//
// It is true because of the pairing function described above:
// e(P, H(m)) = e(pk*G, H(m)) = e(G, pk*H(m)) = e(G, S)
//func Sign(privatekey *PrivateKey, msg string) (*bn256.G1, err error) {
func Sign(privatekey *PrivateKey, msg []byte) (blsSignature []byte, err error) {
	hPoint := hashToG1(msg)
	sig := new(bn256.G1).ScalarMult(hPoint, privatekey.X)

	blsSig := &common.BlsSignature{
		S: sig.Marshal(),
	}

	// convert the signature to json format
	// 将签名格式转换json
	sigContent, err := json.Marshal(blsSig)
	if err != nil {
		return nil, err
	}

	return sigContent, nil

	//	// construct the XuperSignature
	//	// 组装超级签名
	//	xuperSig := &common.XuperSignature{
	//		SigType:    common.BlsSig,
	//		SigContent: sigContent,
	//	}
	//
	//	sig, err := json.Marshal(xuperSig)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return sig, nil
}

func GenerateKeyPair() (*PrivateKey, *PublicKey) {
	sk, pk, _ := bn256.RandomG2(rand.Reader)

	priKey := &PrivateKey{X: sk}
	pubKey := &PublicKey{P: pk}

	return priKey, pubKey
}

//func Verify(publicKey *PublicKey, sig *bn256.G1, msg []byte) (bool, error) {
func Verify(publicKey *PublicKey, sig, msg []byte) (bool, error) {
	signature := new(common.BlsSignature)
	err := json.Unmarshal(sig, signature)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling bls signature [%s]", err)
	}

	sigPointG1 := new(bn256.G1).ScalarBaseMult(big.NewInt(1))
	sigPointG1.Unmarshal(signature.S)

	generatePointG2 := new(bn256.G2).ScalarBaseMult(big.NewInt(1))
	hashPointG1 := hashToG1(msg)

	// e(G, S) = e(S, G)
	lp := bn256.Pair(sigPointG1, generatePointG2)

	// e(P, H(m)) = e(H(m), P)
	rp := bn256.Pair(hashPointG1, publicKey.P)

	// check whether e(G, S) equals e(P, H(m)) or not
	// if sig is valid, then e(P, H(m)) = e(pk*G, H(m)) = e(G, pk*H(m)) = e(G, S)
	isEqual := reflect.DeepEqual(lp.Marshal(), rp.Marshal())

	return isEqual, nil
}

func hashToG1(msg []byte) *bn256.G1 {
	// hash a msg to a point of G1
	k := hash.HashUsingSha256(msg)
	intK := new(big.Int).SetBytes(k)

	return new(bn256.G1).ScalarBaseMult(intK)
}
