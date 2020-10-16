package tss_sign

/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com> 西二旗第一帅
*/

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"fmt"
	//	"log"
	"math/big"

	"github.com/xuperchain/crypto/common/utils"
	"github.com/xuperchain/crypto/core/common"
	"github.com/xuperchain/crypto/core/hash"
)

var (
	InvalidInputParamsError = errors.New("Invalid input params")
	EmptyMessageError       = errors.New("Message to be sign should not be nil")
	NotValidSignatureError  = errors.New("Signature is invalid")
)

//验签算法如下：
//1. 计算：e = hash(C,R,m)
//2. 计算：Rv = sG - eC
//3. 如果Rv == R则返回true，否则返回false
// Because sG = Sum( K(i) + e*X(i)*Coef(i) ) * G = Sum(K(i))*G + Sum(e*X(i)*Coef(i))*G
// = Sum(K(i)*G) + e*Sum(X(i)*Coef(i))*G = R + e*X*G
// = R + eC
// 门限签名的一个核心数学理论是：Sum(e*X(i)*Coef(i)) = X，这个X对应的就是DKG算出的公钥对应的私钥的秘密值
// 因为公钥使用的是所有验证点之和，那么私钥就需要是所有秘密的之和，Coef(i)是为了保证在不泄漏X(i)的情况下，
// 通过计算出自己的系数，最终可以去中心化完成X的计算
func VerifyTssSig(p *ecdsa.PublicKey, signature []byte, message []byte) (bool, error) {
	sig := new(common.TssSignature)
	err := json.Unmarshal(signature, sig)
	if err != nil {
		return false, fmt.Errorf("Failed unmashalling tss signature [%s]", err)
	}

	// sig nil check and sig format check
	if sig == nil || len(sig.R) == 0 || len(sig.S) == 0 {
		return false, NotValidSignatureError
	}

	// empty message
	if len(message) == 0 {
		return false, EmptyMessageError
	}

	curve := p.Curve

	// 计算sG
	lhsX, lhsY := curve.ScalarBaseMult(sig.S)

	// 计算C，converts a point into the uncompressed form specified in section 4.3.6 of ANSI X9.62
	c := elliptic.Marshal(p.Curve, p.X, p.Y)

	// 计算e = HASH(P,R,m)，这里的hash算法选择NIST算法
	hashBytes := hash.HashUsingSha256(utils.BytesCombine(c, sig.R, message))

	// 计算eC,也就是HASH(P,R,m) * C
	rhsX, rhsY := curve.ScalarMult(p.X, p.Y, hashBytes)

	// 计算-eC，如果 eC = (x,y)，则 -eC = (x, -y mod P)
	negativeOne := big.NewInt(-1)
	rhsY = new(big.Int).Mod(new(big.Int).Mul(negativeOne, rhsY), curve.Params().P)

	// 计算Rv = sG - eC
	resX, resY := curve.Add(lhsX, lhsY, rhsX, rhsY)

	// 原始签名中的R
	rX, rY := elliptic.Unmarshal(curve, sig.R)

	// 对比签名是否一致
	if resX.Cmp(rX) == 0 && resY.Cmp(rY) == 0 {
		return true, nil
	}

	return false, nil
}
