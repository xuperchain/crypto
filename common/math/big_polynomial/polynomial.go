package big_polynomial

/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com> 西二旗第一帅
*/

import (
	//	"log"
	//	"math"
	"math/big"
	//	"fmt"

	"github.com/xuperchain/crypto/common/math/rand"
	//	"github.com/xuperchain/crypto/core/hdwallet/rand"
)

//const (
//	// A big prime which is used for Galois Field computing
//	primeStr = "24815323469403931728221172233738523533528335161133543380459461440894543366372904768334987263999999999999999999663"
//)
//
//var (
//	prime, _ = big.NewInt(0).SetString(pc.primeStr, 10)
//
////	prime = elliptic.P256().Params().N
//)

type PolynomialClient struct {
	// A big prime which is used for Galois Field computing
	prime *big.Int
}

func New(prime *big.Int) *PolynomialClient {
	pc := new(PolynomialClient)
	pc.prime = prime

	return pc
}

// make a random polynomials F(x) of Degree [degree], and the const(X-Intercept) is [intercept]
// 给定最高次方和x截距，生成一个系数随机的多项式
func (pc *PolynomialClient) RandomGenerate(degree int, secret []byte) ([]*big.Int, error) {
	// 字节数组转big int
	intercept := big.NewInt(0).SetBytes(secret)

	// 多项式参数格式是次方数+1（代表常数）
	result := make([]*big.Int, degree+1)

	// 多项式的常数项就是x截距
	// 多个bytes组成一个bigint，作为多项式的系数
	coefficientFactor := 32
	byteSlice := make([]byte, coefficientFactor)
	index := 0
	result[index] = intercept

	// 生成非最高次方位的随机参数
	if degree > 1 {
		randomBytes, err := rand.GenerateSeedWithStrengthAndKeyLen(rand.KeyStrengthHard, coefficientFactor*(degree-1))
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(randomBytes); i += coefficientFactor {
			byteSlice = randomBytes[i : i+coefficientFactor]
			//			log.Printf("byteSlice is %v", byteSlice)
			result[index+1] = big.NewInt(0).SetBytes(byteSlice)
			index++
		}
	}

	// This coefficient can't be zero, otherwise it will be a polynomial,
	// the degree of which is [degree-1] other than [degree]
	// 生成最高次方位的随机参数，该值不能为0，否则最高次方会退化为次一级
	for {
		randomBytes, err := rand.GenerateSeedWithStrengthAndKeyLen(rand.KeyStrengthHard, coefficientFactor)
		if err != nil {
			return nil, err
		}

		highestDegreeCoefficient := big.NewInt(0).SetBytes(randomBytes)
		if highestDegreeCoefficient != big.NewInt(0) {
			result[degree] = highestDegreeCoefficient
			return result, nil
		}
	}
}

// Given the specified value, get the compution result of the polynomial
// 给出指定x值，计算出指定多项式f(x)的值
func (pc *PolynomialClient) Evaluate(polynomialCoefficients []*big.Int, specifiedValue *big.Int) *big.Int {
	//	log.Printf("polynomialCoefficients is: %v and specifiedValue is %v", polynomialCoefficients, specifiedValue)
	degree := len(polynomialCoefficients) - 1

	// 注意这里要用set，否则会出现上层业务逻辑的指针重复使用的问题
	result := big.NewInt(0).Set(polynomialCoefficients[degree])

	for i := degree - 1; i >= 0; i-- {
		result = result.Mul(result, specifiedValue)
		result = result.Add(result, polynomialCoefficients[i])
	}

	return result
}

// 对2个多项式进行加法操作
func (pc *PolynomialClient) Add(a []*big.Int, b []*big.Int) []*big.Int {
	degree := len(a)
	c := make([]*big.Int, degree)

	// 初始化big int数组
	for i, _ := range c {
		c[i] = big.NewInt(0)
	}

	for i := 0; i < degree; i++ {
		c[i] = a[i].Add(a[i], b[i])

		// 域运算
		c[i] = big.NewInt(0).Mod(c[i], pc.prime)
	}

	return c
}

// 对2个多项式进行乘法操作
func (pc *PolynomialClient) Multiply(a []*big.Int, b []*big.Int) []*big.Int {
	degA := len(a)
	degB := len(b)
	result := make([]*big.Int, degA+degB-1)

	// 初始化big int数组
	for i, _ := range result {
		result[i] = big.NewInt(0)
	}

	for i := 0; i < degA; i++ {
		for j := 0; j < degB; j++ {
			temp := a[i].Mul(a[i], b[j])
			//			log.Printf("temp is %v", temp)
			//			log.Printf("result[i+j] is %v", result[i+j])
			result[i+j] = result[i+j].Add(result[i+j], temp)
		}
	}

	return result
}

// 将1个多项式与指定系数k进行乘法操作
func (pc *PolynomialClient) Scale(a []*big.Int, k *big.Int) []*big.Int {
	b := make([]*big.Int, len(a))

	for i := 0; i < len(a); i++ {
		b[i] = a[i].Mul(a[i], k)

		// 域运算
		b[i] = big.NewInt(0).Mod(b[i], pc.prime)
	}

	return b
}

// 获取拉格朗日基本多项式（插值基函数）
func (pc *PolynomialClient) GetLagrangeBasePolynomial(xs []*big.Int, xpos int) []*big.Int {
	var poly []*big.Int
	poly = append(poly, big.NewInt(1))

	// 分母
	denominator := big.NewInt(1)

	for i := 0; i < len(xs); i++ {
		if i != xpos {
			currentTerm := make([]*big.Int, 2)
			currentTerm[0] = big.NewInt(1)
			//			currentTerm[1] = -xs[i]
			currentTerm[1] = big.NewInt(0).Sub(big.NewInt(0), xs[i])
			//			denominator *= xs[xpos] - xs[i]
			denominator = denominator.Mul(denominator, big.NewInt(0).Sub(xs[xpos], xs[i]))
			poly = pc.Multiply(poly, currentTerm)
		}
	}

	//	log.Printf("getLagrangeBasePolynomial poly is: %v and denominator is %v", poly, denominator)
	inverser := big.NewInt(0).ModInverse(denominator, pc.prime)
	//	log.Printf("scale factor is: %v and denominator is: %v", inverser, denominator)

	// 校验是否在有限域上的取逆操作是正确的
	//	tmp := big.NewInt(0).Mul(denominator, inverser)
	//	tmp = big.NewInt(0).Mod(tmp, pc.prime)
	//	log.Printf("inverser check result is: %v", tmp)

	return pc.Scale(poly, inverser)
}

// 利用Lagrange Polynomial Interpolation Formula，通过给定坐标点集合来计算多项式
func (pc *PolynomialClient) GetPolynomialByPoints(points map[int]*big.Int) []*big.Int {
	degree := len(points)
	bases := make([][]*big.Int, degree)
	result := make([]*big.Int, degree)

	// 初始化big int数组
	for i, _ := range result {
		result[i] = big.NewInt(0)
	}

	var xs []*big.Int
	var ys []*big.Int

	for k, v := range points {
		xs = append(xs, big.NewInt(int64(k)))
		ys = append(ys, v)
	}

	for i := 0; i < degree; i++ {
		bases[i] = pc.GetLagrangeBasePolynomial(xs, i)
	}

	for i := 0; i < degree; i++ {
		//		log.Printf("The coefficients of the lagrange base polynomial[%d] is: %v", i, bases[i])
		result = pc.Add(result, pc.Scale(bases[i], ys[i]))
	}

	return result
}
