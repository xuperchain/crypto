/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com>
*/

package schnorr_sign_new

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/xuperchain/crypto/core/common"
	"github.com/xuperchain/crypto/core/hash"
)

var (
	GenerateSignatureError = errors.New("Failed to generate the schnorr signature, s = 0 happened.")
	EmptyMessageError      = errors.New("The message to be signed should not be empty")
)

// Schnorr signatures use a particular function, defined as:
// H'(m, s, e) = H(m || s * G - e * P)
//
// H is a hash function, for instance SHA256 or SM3.
// s and e are 2 numbers forming the signature itself.
// m is the message to sign.
// P is the public key.
//
// To verify the signature, check that the result of H'(m, s, e) is equal to e.
// Which means that: H(m || s * G - e * P) = e
//
// It's impossible for the others to find such a pair of (s, e) but the signer himself.
// This is because: P = x * G
// So the signer is able to get this equation: H(m || s * G - e * x * G) = e = H(m || (s - e * x) * G)
// It can be considered as:  H(m || k * G) = e, where k = s - e * x
//
// This is the original process:
// 1. Choose a random number k
// 2. Compute e = H(m || k * G)
// 3. Because k = s - e * x, k and x (the key factor of the private key) are already known, we can compute s
// 4. Now we get the SchnorrSignature (e, s)
//
// Note that there is a potential risk for the private key, which also exists in the ECDSA algorithm:
// "The number k must be random enough."
// If not, say the same k has been used twice or the second k can be predicted by the first k,
// the attacker will be able to retrieve the private key (x)
// This is because:
// 1. If the same k has been used twice:
//    k = s0 - e0 * x = s1 - e1 * x
// The attacker knows: x = (s0 - s1) / (e0 - e1)
//
// 2. If the second k1 can be predicted by the first k0:
//    k0 = s0 - e0 * x
//    k1 = s1 - e1 * x
// The attacker knows: x = (k1 - k0 + s0 - s1) / (e0 - e1)
//
// So the final process is:
// 1. Compute k = H(m || x)
//    This makes k unpredictable for anyone who do not know x,
//    therefor it's impossible for the attacker to retrive x by breaking the random number generator of the system,
//    which has happend in the Sony PlayStation 3 firmware attack.
// 2. Compute e = H(m || k * G)
// 3. Because k = s - e * x, k and x (the key factor of the private key) are already known,
//    we can compute s = k + e * x
// 4. Now we get the SchnorrSignature (e, s)
func Sign(privateKey *ecdsa.PrivateKey, message []byte) (schnorrSignature []byte, err error) {
	if privateKey == nil {
		return nil, fmt.Errorf("Invalid privateKey. PrivateKey must not be nil.")
	}

	// 1. Compute k = H(m || x)
	k := hash.HashUsingSha256(append(message, privateKey.D.Bytes()...))

	// 2. Compute e = H(m || k * G)
	// 2.1 compute k * G
	curve := privateKey.Curve
	x, y := curve.ScalarBaseMult(k)
	// 2.2 compute H(m || k * G)
	e := hash.HashUsingSha256(append(message, elliptic.Marshal(curve, x, y)...))

	// 3. k = s - e * x, so we can compute s = k + e * x
	intK := new(big.Int).SetBytes(k)
	intE := new(big.Int).SetBytes(e)

	intS, err := ComputeSByKEX(curve, intK, intE, privateKey.D)
	if err != nil {
		return nil, GenerateSignatureError
	}

	// generate the schnorr signature：(sum(S), R)
	// 生成Schnorr签名：(sum(S), R)
	schnorrSig := &common.SchnorrSignature{
		E: intE,
		S: intS,
	}
	// convert the signature to json format
	// 将签名格式转换json
	sigContent, err := json.Marshal(schnorrSig)
	if err != nil {
		return nil, err
	}

	return sigContent, nil

	// construct the XuperSignature
	// 组装超级签名
	xuperSig := &common.XuperSignature{
		SigType:    common.Schnorr,
		SigContent: sigContent,
	}

	sig, err := json.Marshal(xuperSig)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

// Compute s = k + e*x
func ComputeSByKEX(curve elliptic.Curve, k, e, x *big.Int) (*big.Int, error) {
	intS := new(big.Int).Add(k, new(big.Int).Mul(e, x))

	return intS, nil
}

// In order to verify the signature, only need to check the equation:
// H'(m, s, e) = H(m || s * G - e * P) = e
// i.e. whether e is equal to H(m || s * G - e * P)
func Verify(publicKey *ecdsa.PublicKey, sig []byte, message []byte) (valid bool, err error) {
	signature := new(common.SchnorrSignature)
	err = json.Unmarshal(sig, signature)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling schnorr signature [%s]", err)
	}

	// 1. compute h(m|| s * G - e * P)
	// 1.1 compute s * G
	curve := publicKey.Curve
	x1, y1 := curve.ScalarBaseMult(signature.S.Bytes())

	// 1.2 compute e * P
	x2, y2 := curve.ScalarMult(publicKey.X, publicKey.Y, signature.E.Bytes())

	// 1.3 计算-(e * P)，如果 e * P = (x,y)，则 -(e * P) = (x, -y mod P)
	negativeOne := big.NewInt(-1)
	y2 = new(big.Int).Mod(new(big.Int).Mul(negativeOne, y2), curve.Params().P)

	// 1.4 compute s * G - e * P
	x, y := curve.Add(x1, y1, x2, y2)

	e := hash.HashUsingSha256(append(message, elliptic.Marshal(curve, x, y)...))

	intE := new(big.Int).SetBytes(e)

	// 2. check the equation
	//	return bytes.Equal(e, signature.E.Bytes()), nil
	if intE.Cmp(signature.E) != 0 {
		return false, nil
	}
	return true, nil
}
