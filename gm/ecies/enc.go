/*
Copyright Baidu Inc. All Rights Reserved.

jingbo@baidu.com
*/
package ecies

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/xuperchain/crypto/gm/config"
	"github.com/xuperchain/crypto/gm/gmsm/sm2"
)

func Encrypt(k *ecdsa.PublicKey, msg []byte) (cypherText []byte, err error) {
	// 判断是否是国密标准的公钥
	isGmCurve := checkKeyCurve(k)
	if isGmCurve == false {
		return nil, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}

	key := new(sm2.PublicKey)
	//	key := &sm2.PrivateKey{}
	key.Curve = sm2.P256Sm2() // elliptic.P256()
	key.X = k.X
	key.Y = k.Y

	cypherText, err = sm2.Encrypt(key, msg)
	return cypherText, err

}

// 判断是否是国密标准的公钥
func checkKeyCurve(k *ecdsa.PublicKey) bool {
	if k.X == nil || k.Y == nil {
		return false
	}

	switch k.Params().Name {
	case config.CurveNist: // NIST
		return false
	case config.CurveGm: // 国密
		return true
	default: // 不支持的密码学类型
		return false
	}
}

func Decrypt(k *ecdsa.PrivateKey, cypherText []byte) (msg []byte, err error) {
	// 判断是否是国密标准的私钥
	isGmCurve := checkKeyCurve(&k.PublicKey)
	if isGmCurve == false {
		return nil, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}
	if k.D == nil {
		return nil, fmt.Errorf("Param D cannot be nil.")
	}

	key := new(sm2.PrivateKey)
	//	key := &sm2.PrivateKey{}
	key.PublicKey.Curve = sm2.P256Sm2() // elliptic.P256()
	key.X = k.X
	key.Y = k.Y
	key.D = k.D

	msg, err = sm2.Decrypt(key, cypherText)
	return msg, err
}
