package complex_secret_share

/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com> 西二旗第一帅
*/

import (
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"log"
	"math/big"

	"github.com/xuperchain/crypto/common/math/ecc"

	polynomial "github.com/xuperchain/crypto/common/math/big_polynomial"
	//	polynomial "github.com/xuperchain/crypto/core/secret_share/big_polynomial"
)

var (
	InvaildTotalShareNumberError = errors.New("The totalShareNumber must be greater than one.")
	InvaildShareNumberError      = errors.New("The minimumShareNumber must be smaller than the totalShareNumber.")
)

// Shamir's Secret Sharing algorithm, can be considered as:
// A way to split a secret to W shares, the secret can only be retrieved if more than T(T <= W) shares are combined together.
//
// This is the split process:
// 1. Encode the secret to a number S
// 2. Choose a lot of random numbers as coefficients, in order to make a random polynomials F(x) of degree T-1,
//		 the variable is X, the const(x-intercept) is S
// 3. For this polynomial, Give x diffent values, for example, x++ each time, then compute y = F(x)
// 4. So we get W shares, which are (x, y) pairs
// 5. Now encode each pair to a byte slice
func ComplexSecretSplit(totalShareNumber, minimumShareNumber int, secret []byte, curve elliptic.Curve) (shares map[int]*big.Int, err error) {
	poly, err := ComplexSecretToPolynomial(totalShareNumber, minimumShareNumber, secret, curve)
	if err != nil {
		return nil, err
	}

	//	log.Printf("The asc order coefficients of the polynomial is: %v", poly)

	polynomialClient := polynomial.New(curve.Params().N)

	// Evaluate the polynomial for several times, in order to get all the shares.
	shares = make(map[int]*big.Int, totalShareNumber)
	for x := 1; x <= totalShareNumber; x++ {
		shares[x] = polynomialClient.Evaluate(poly, big.NewInt(int64(x)))
	}
	//	log.Printf("shares is: %v", shares)

	return shares, nil
}

func ComplexSecretSplitWithVerifyPoints(totalShareNumber, minimumShareNumber int, secret []byte, curve elliptic.Curve) (shares map[int]*big.Int, points []*ecc.Point, err error) {
	poly, err := ComplexSecretToPolynomial(totalShareNumber, minimumShareNumber, secret, curve)
	if err != nil {
		return nil, nil, err
	}

	//	log.Printf("The asc order coefficients of the polynomial is: %v", poly)

	//	points = make([]*ecc.Point, len(poly))
	//	var points []*ecc.Point
	for _, coefficient := range poly {
		x, y := elliptic.P256().ScalarBaseMult(coefficient.Bytes())
		//		point, err := ecc.NewPoint(elliptic.P256(), x, y)
		point, err := ecc.NewPoint(curve, x, y)
		//		log.Printf("coefficient is %v, point is: %v", coefficient, point)
		if err != nil {
			return nil, nil, err
		}
		points = append(points, point)
	}
	jsonPoints, _ := json.Marshal(points)
	log.Printf("json points is: %s", jsonPoints)

	polynomialClient := polynomial.New(curve.Params().N)

	// Evaluate the polynomial for several times, in order to get all the shares.
	shares = make(map[int]*big.Int, totalShareNumber)
	for x := 1; x <= totalShareNumber; x++ {
		shares[x] = polynomialClient.Evaluate(poly, big.NewInt(int64(x)))
	}
	//	log.Printf("shares is: %v", shares)

	return shares, points, nil
}

func ComplexSecretToPolynomial(totalShareNumber, minimumShareNumber int, secret []byte, curve elliptic.Curve) ([]*big.Int, error) {
	// Check the parameters
	if totalShareNumber < 2 {
		return nil, InvaildTotalShareNumberError
	}

	if minimumShareNumber > totalShareNumber {
		return nil, InvaildShareNumberError
	}

	polynomialClient := polynomial.New(curve.Params().N)

	poly, err := polynomialClient.RandomGenerate(minimumShareNumber-1, secret)
	if err != nil {
		return nil, err
	}

	return poly, nil
}

// 为产生本地秘密的私钥碎片做准备，通过目标多项式生成验证点
func GetVerifyPointByPolynomial(poly []*big.Int, curve elliptic.Curve) (*ecc.Point, error) {
	x, y := elliptic.P256().ScalarBaseMult(poly[0].Bytes())
	point, err := ecc.NewPoint(curve, x, y)
	if err != nil {
		return nil, err
	}

	return point, nil
}

// 为产生本地秘密的私钥碎片做准备，通过目标多项式和节点index生成对应的碎片
func GetSpecifiedSecretShareByPolynomial(poly []*big.Int, index *big.Int, curve elliptic.Curve) *big.Int {
	polynomialClient := polynomial.New(curve.Params().N)

	// Evaluate the polynomial with the specified index, in order to get the shares.
	share := polynomialClient.Evaluate(poly, index)

	return share
}
