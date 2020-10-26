package utils

import (
	"crypto/ecdsa"

	"github.com/xuperchain/crypto/common/math/curve"
)

func ChangePrivCurveToS256k1(key *ecdsa.PrivateKey) *ecdsa.PrivateKey {
	s256k1Curve := curve.Secp256k1()

	priv := &ecdsa.PrivateKey{}
	priv.PublicKey.Curve = s256k1Curve
	priv.D = key.D

	// 重新计算公钥，因为基点变了
	priv.PublicKey.X, priv.PublicKey.Y = s256k1Curve.ScalarBaseMult(key.D.Bytes())

	return priv
}
