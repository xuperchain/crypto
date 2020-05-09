package sign

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"

	"github.com/xuperchain/crypto/gm/common"
	"github.com/xuperchain/crypto/gm/config"
	"github.com/xuperchain/crypto/gm/gmsm/sm2"
)

func SignECDSA(k *ecdsa.PrivateKey, msg []byte) (signature []byte, err error) {
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

	r, s, err := sm2.Sign(key, msg)
	if err != nil {
		return nil, fmt.Errorf("Failed to sign the msg [%s]", err)
	}
	return MarshalECDSASignature(r, s)

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

func SignV2ECDSA(k *ecdsa.PrivateKey, msg []byte) (signature []byte, err error) {
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

	r, s, err := sm2.Sign(key, msg)
	if err != nil {
		return nil, fmt.Errorf("Failed to sign the msg [%s]", err)
	}

	// 生成ECDSA签名：(sum(S), R)
	ecdsaSig := &common.ECDSASignature{
		R: r,
		S: s,
	}

	// 生成超级签名
	// 转换json
	sigContent, err := json.Marshal(ecdsaSig)
	if err != nil {
		return nil, err
	}

	xuperSig := &common.XuperSignature{
		SigType:    common.ECDSA,
		SigContent: sigContent,
	}

	sig, err := json.Marshal(xuperSig)
	if err != nil {
		return nil, err
	}

	return sig, nil
}

func VerifyV2ECDSA(k *ecdsa.PublicKey, sig, msg []byte) (valid bool, err error) {
	// 判断是否是国密标准的公钥
	isGmCurve := checkKeyCurve(k)
	if isGmCurve == false {
		return false, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}

	signature := new(common.ECDSASignature)
	err = json.Unmarshal(sig, signature)
	if err != nil {
		return false, fmt.Errorf("Failed to unmarshal the ecdsa signature [%s]", err)
	}

	key := new(sm2.PublicKey)
	key.Curve = sm2.P256Sm2() // elliptic.P256()
	key.X = k.X
	key.Y = k.Y

	return sm2.Verify(key, msg, signature.R, signature.S), nil
}

func VerifyECDSA(k *ecdsa.PublicKey, sig, msg []byte) (valid bool, err error) {
	// 判断是否是国密标准的公钥
	isGmCurve := checkKeyCurve(k)
	if isGmCurve == false {
		return false, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}

	r, s, err := UnmarshalECDSASignature(sig)
	if err != nil {
		return false, fmt.Errorf("Failed to unmarshal the ecdsa signature [%s]", err)
	}

	key := new(sm2.PublicKey)
	key.Curve = sm2.P256Sm2() // elliptic.P256()
	key.X = k.X
	key.Y = k.Y

	return sm2.Verify(key, msg, r, s), nil
}
