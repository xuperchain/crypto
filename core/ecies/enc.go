/*
Copyright Baidu Inc. All Rights Reserved.

jingbo@baidu.com
*/
package ecies

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"

	"github.com/xuperchain/crypto/core/config"
	libecies "github.com/xuperchain/crypto/core/ecies/libecies"
)

func Encrypt(k *ecdsa.PublicKey, msg []byte) (cypherText []byte, err error) {
	// 判断是否是NIST标准的公钥
	isNistCurve := checkKeyCurve(k)
	if isNistCurve == false {
		return nil, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}

	pub := libecies.ImportECDSAPublic(k)

	ct, err := libecies.Encrypt(rand.Reader, pub, msg, nil, nil)
	if err != nil {
		return nil, err
	}

	return ct, nil
}

// 判断是否是NIST标准的公钥
func checkKeyCurve(k *ecdsa.PublicKey) bool {
	if k.X == nil || k.Y == nil {
		return false
	}

	switch k.Params().Name {
	case config.CurveNist: // NIST
		return true
	case config.CurveGm: // 国密
		return false
	default: // 不支持的密码学类型
		return false
	}
}

func Decrypt(k *ecdsa.PrivateKey, cypherText []byte) (msg []byte, err error) {
	// 判断是否是NIST标准的私钥
	isNistCurve := checkKeyCurve(&k.PublicKey)
	if isNistCurve == false {
		return nil, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}
	if k.D == nil {
		return nil, fmt.Errorf("Param D cannot be nil.")
	}

	prv := libecies.ImportECDSA(k)

	pt, err := prv.Decrypt(rand.Reader, cypherText, nil, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return pt, nil
}
