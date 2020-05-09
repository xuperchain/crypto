package sign

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"

	"github.com/xuperchain/crypto/core/common"
	"github.com/xuperchain/crypto/core/config"
)

func SignECDSA(k *ecdsa.PrivateKey, msg []byte) (signature []byte, err error) {
	// 判断是否是NIST标准的私钥
	isNistCurve := checkKeyCurve(&k.PublicKey)
	if isNistCurve == false {
		return nil, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}
	if k.D == nil {
		return nil, fmt.Errorf("Param D cannot be nil.")
	}

	r, s, err := ecdsa.Sign(rand.Reader, k, msg)
	if err != nil {
		return nil, err
	}

	return MarshalECDSASignature(r, s)
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

func SignV2ECDSA(k *ecdsa.PrivateKey, msg []byte) (signature []byte, err error) {
	// 判断是否是NIST标准的私钥
	isNistCurve := checkKeyCurve(&k.PublicKey)
	if isNistCurve == false {
		return nil, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}
	if k.D == nil {
		return nil, fmt.Errorf("Param D cannot be nil.")
	}

	r, s, err := ecdsa.Sign(rand.Reader, k, msg)
	if err != nil {
		return nil, err
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
	// 判断是否是NIST标准的公钥
	isNistCurve := checkKeyCurve(k)
	if isNistCurve == false {
		return false, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}

	signature := new(common.ECDSASignature)
	err = json.Unmarshal(sig, signature)
	if err != nil {
		return false, fmt.Errorf("Failed to unmarshal the ecdsa signature [%s]", err)
	}

	return ecdsa.Verify(k, msg, signature.R, signature.S), nil
}

func VerifyECDSA(k *ecdsa.PublicKey, sig, msg []byte) (valid bool, err error) {
	// 判断是否是NIST标准的公钥
	isNistCurve := checkKeyCurve(k)
	if isNistCurve == false {
		return false, fmt.Errorf("This cryptography curve[%s] has not been supported yet.", k.Params().Name)
	}

	r, s, err := UnmarshalECDSASignature(sig)
	if err != nil {
		return false, fmt.Errorf("Failed to unmarshal the ecdsa signature [%s]", err)
	}

	return ecdsa.Verify(k, msg, r, s), nil
}
