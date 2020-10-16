package tss_sign

/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com> 西二旗第一帅
*/

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/json"
	//	"log"
	"math/big"

	"github.com/xuperchain/crypto/common/utils"
	"github.com/xuperchain/crypto/core/common"
	"github.com/xuperchain/crypto/core/hash"
	"github.com/xuperchain/crypto/core/hdwallet/rand"

	polynomial "github.com/xuperchain/crypto/common/math/big_polynomial"
)

// 生成默认随机数Ki
func GetRandom32Bytes() ([]byte, error) {
	randomBytes, err := rand.GenerateSeedWithStrengthAndKeyLen(rand.KeyStrengthHard, rand.KeyLengthInt32)
	if err != nil {
		return nil, err
	}

	return randomBytes, nil
}

// 计算：Ri = Ki*G
func GetRiUsingRandomBytes(key *ecdsa.PublicKey, k []byte) []byte {
	curve := key.Curve

	// 计算K*G
	x, y := curve.ScalarBaseMult(k)

	// 计算R，converts a point into the uncompressed form specified in section 4.3.6 of ANSI X9.62
	r := elliptic.Marshal(curve, x, y)

	return r
}

// 计算：R = k1*G + k2*G + ... + kn*G
func GetRUsingAllRi(key *ecdsa.PublicKey, arrayOfRi [][]byte) []byte {
	num := len(arrayOfRi)
	curve := key.Curve
	x, y := big.NewInt(0), big.NewInt(0)
	for i := 0; i < num; i++ {
		// Unmarshal converts a point, serialized by Marshal, into an x, y pair.
		// It is an error if the point is not in uncompressed form or is not on the curve.
		// On error, x = nil.
		x1, y1 := elliptic.Unmarshal(curve, arrayOfRi[i])

		// 计算k1*G + k2*G + ...
		x, y = curve.Add(x, y, x1, y1)
	}
	// 计算R，converts a point into the uncompressed form specified in section 4.3.6 of ANSI X9.62
	r := elliptic.Marshal(curve, x, y)

	return r
}

// 获取Si中Xi的系数，不同的Xi的系数不一样，所有的Xi乘以自己的系数，
// 再加起来之和，就是最终的通过所有点(x，sum(y))的多项式f(x)
// f(x)当x=0时的值，就是所有参与方的所保管的秘密之和
// 为了获得这个系数，需要通过拉格朗日插值公式Lagrange interpolation formula
// 来获得插值基函数Lagrange base polynomial，
// 但是我们并不需要获得每个基多项式，我们要的只是每个基多项式当x=0时的值
func getCoefForXi(xs []*big.Int, xpos int, curve elliptic.Curve) *big.Int {
	polynomialClient := polynomial.New(curve.Params().N)

	poly := polynomialClient.GetLagrangeBasePolynomial(xs, xpos)
	//	log.Printf("coef is [%d]:", poly[len(poly)-1])

	// 当xpos为0时，常数项的系数
	coef := poly[len(poly)-1]

	return coef
}

// 获取Si中的(coefi*Xi)，注意，在门限签名中，S(i) = K(i) + HASH(C,R,m) * Coef(i) * X(i)
// 每个实际参与节点再次计算自己的系数Coef(i)，为下一步的S(i)计算做准备
// indexSet是指所有实际参与节点的index所组成的集合
// localIndexPos是本节点在indexSet中的位置
// key是在DKG过程中，自己计算出的私钥
//func GetXiWithcoef(coef, xi *big.Int) *big.Int {
func GetXiWithcoef(xs []*big.Int, xpos int, key *ecdsa.PrivateKey) *big.Int {
	xi := key.D
	curve := key.Curve

	coef := getCoefForXi(xs, xpos, curve)

	coef = new(big.Int).Mul(coef, xi)

	return coef
}

// 计算 s(i) = k(i) + HASH(C,R,m) * x(i) * coef(i)
// x代表大数D，也就是私钥的关键参数
func GetSiUsingKCRMWithCoef(k []byte, c []byte, r []byte, message []byte, coef *big.Int) []byte {
	// 计算HASH(P,R,m)，这里的hash算法选择NIST算法
	hashBytes := hash.HashUsingSha256(utils.BytesCombine(c, r, message))

	// 计算HASH(P,R,m) * xi * coef
	tmpResult := new(big.Int).Mul(new(big.Int).SetBytes(hashBytes), coef)

	// 计算ki + HASH(P,R,m) * xi * coef
	s := new(big.Int).Add(new(big.Int).SetBytes(k), tmpResult)

	return s.Bytes()
}

// 计算 s(i) = HASH(C,R,m) * x(i) * coef(i)
// x代表大数D，也就是私钥的关键参数
func GetSiUsingKCRMWithCoefNoKi(c []byte, r []byte, message []byte, coef *big.Int) []byte {
	// 计算HASH(P,R,m)，这里的hash算法选择NIST算法
	hashBytes := hash.HashUsingSha256(utils.BytesCombine(c, r, message))

	// 计算HASH(P,R,m) * xi * coef
	s := new(big.Int).Mul(new(big.Int).SetBytes(hashBytes), coef)

	return s.Bytes()
}

// 计算：S = sum(si)
func GetSUsingAllSi(arrayOfSi [][]byte) []byte {
	num := len(arrayOfSi)
	s := big.NewInt(0)
	for i := 0; i < num; i++ {
		// 计算s1 + s2 + ... + sn
		s = s.Add(s, new(big.Int).SetBytes(arrayOfSi[i]))
	}

	return s.Bytes()
}

//生成门限签名的流程如下：
//1. 各方分别生成自己的随机数Ki(K1, K2, ..., Kn) --- func getRandomBytes() ([]byte, error)
//      Compute k = H(m || x), m is the msg to be signed and x is the private key of the node.
//    	This makes k unpredictable for anyone who do not know x,
//    	therefor it's impossible for the attacker to retrive x by breaking the random number generator of the system,
//   	which has happend in the Sony PlayStation 3 firmware attack.
//		不再使用临时随机数，而改用H(m || x)来计算k
//2. 各方计算自己的 Ri = Ki*G，G代表基点 --- func getRiUsingRandomBytes(key *ecdsa.PublicKey, k []byte) []byte
//3. 发起者收集Ri，计算：R = sum(Ri) --- func getRUsingAllRi(key *ecdsa.PublicKey, arrayOfRi [][]byte) []byte
//4. 发起者收集验证节点，计算公共公钥：C = VP(1) + VP(2) + ... + VP(i)
//5. 各方根据自己的index值，和本次计算所有参与方的index集合，计算出自己的Coef
//6. 各方计算自己的S(i)：S(i) = K(i) + HASH(C,R,m) * Coef(i) * X(i)，X代表私钥中的参数大数D
// --- func getSiUsingKCRM(key *ecdsa.PrivateKey, k []byte, c []byte, r []byte, message []byte) []byte
//7. 发起者收集Si，生成门限签名：(s1 + s2 + ... + sn, R)
// GenerateTssSignSignature生成对特定消息的门限签名，所有参与签名的私钥必须使用同一条椭圆曲线
func GenerateTssSignSignature(s []byte, r []byte) ([]byte, error) {
	// 生成门限签名：(sum(S), R)
	tssSig := &common.TssSignature{
		S: s,
		R: r,
	}

	// 生成超级签名
	// 转换json
	sigContent, err := json.Marshal(tssSig)
	if err != nil {
		return nil, err
	}

	xuperSig := &common.XuperSignature{
		SigType:    common.TssSig,
		SigContent: sigContent,
	}

	//	log.Printf("xuperSig before marshal: %s", xuperSig)

	sig, err := json.Marshal(xuperSig)
	if err != nil {
		return nil, err
	}

	return sig, nil

}
