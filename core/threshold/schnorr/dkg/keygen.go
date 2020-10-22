package dkg

/*
Copyright Baidu Inc. All Rights Reserved.

<jingbo@baidu.com> 西二旗第一帅
*/

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"errors"
	//	"log"
	"math/big"

	"github.com/xuperchain/crypto/common/math/ecc"

	walletRand "github.com/xuperchain/crypto/core/hdwallet/rand"
	secret_share "github.com/xuperchain/crypto/core/secret_share/complex_secret_share"
)

var (
	zero = big.NewInt(0)
)

type (
	PartnerShares struct {
		//		Index int
		PartnerInfo *PartnerPublic
		// key: partner index，也就是x坐标, value: 实际数值，也就是y坐标
		Shares map[int]*big.Int
		//		VerifyPoints map[*big.Int]*big.Int
		VerifyPoints []*ecc.Point
	}

	PartnerPublic struct {
		Index        int
		IndentityKey *big.Int
	}

	PartnerPrivate struct {
		PublicInfo *PartnerPublic
		Xi         *big.Int
		// public keys (P(i) = secret(i)*G for each Partner(i))
		publicKeyPoints []*ecc.Point
	}
)

// 为产生本地秘密的私钥碎片做准备，预先生成好一个目标多项式
// minimumShareNumber可以理解为threshold
func GetPolynomialForSecretShareGenerate(totalShareNumber, minimumShareNumber int) ([]*big.Int, error) {
	// 1. calculate "partial" key share
	secret, err := walletRand.GenerateEntropy(120)
	if err != nil {
		return nil, err
	}

	curve := elliptic.P256()

	// 2. compute the polynomial for generating the shares and the verify points in the future
	poly, err := secret_share.ComplexSecretToPolynomial(totalShareNumber, minimumShareNumber, secret, curve)
	if err != nil {
		return nil, err
	}

	return poly, nil
}

// 为产生本地秘密的私钥碎片做准备，通过目标多项式生成验证点
func GetVerifyPointByPolynomial(poly []*big.Int) (*ecc.Point, error) {
	curve := elliptic.P256()

	point, err := secret_share.GetVerifyPointByPolynomial(poly, curve)
	if err != nil {
		return nil, err
	}

	return point, nil
}

// 为产生本地秘密的私钥碎片做准备，通过目标多项式和节点index生成对应的碎片
func GetSpecifiedSecretShareByPolynomial(poly []*big.Int, index *big.Int) *big.Int {
	curve := elliptic.P256()

	share := secret_share.GetSpecifiedSecretShareByPolynomial(poly, index, curve)

	return share
}

// 产生本地秘密的私钥碎片，可以把每个碎片理解为一个坐标点。 key: partner index，也就是x坐标, value: 实际数值，也就是y坐标
// minimumShareNumber可以理解为threshold
func LocalSecretShareGenerateWithVerifyPoints(totalShareNumber, minimumShareNumber int) (shares map[int]*big.Int, points []*ecc.Point, err error) {
	// 1. calculate "partial" key share
	secret, err := walletRand.GenerateEntropy(120)
	if err != nil {
		return nil, nil, err
	}

	curve := elliptic.P256()

	// 2. compute the shares and the verify points
	shares, points, err = secret_share.ComplexSecretSplitWithVerifyPoints(totalShareNumber, minimumShareNumber, secret, curve)
	if err != nil {
		return nil, nil, err
	}

	return shares, points, nil
}

// 每个参与节点根据所收集的所有的与自己相关的碎片(自己的Index是X值，收集所有该X值对应的Y值)，
// 来计算出自己的本地私钥xi(该X值对应的Y值之和)，这是一个关键秘密信息
//func LocalPrivateKeyGenerate(shares map[int]*big.Int) *ecdsa.PrivateKey {
func LocalPrivateKeyGenerate(shares []*big.Int) *ecdsa.PrivateKey {
	localXi := big.NewInt(0)

	for _, value := range shares {
		localXi = new(big.Int).Add(localXi, value)
	}

	curve := elliptic.P256()

	localXi = new(big.Int).Mod(localXi, curve.Params().N)

	localPrivateKey := new(ecdsa.PrivateKey)
	localPrivateKey.PublicKey.Curve = curve
	localPrivateKey.D = localXi
	localPrivateKey.PublicKey.X, localPrivateKey.PublicKey.Y = curve.ScalarBaseMult(localXi.Bytes())

	return localPrivateKey
}

// 根据所有潜在节点发布的验证点集合中的第一个元素（也就是秘密值的验证点），计算出公钥
func PublicKeyGenerate(verifyPoints []*ecc.Point) (*ecdsa.PublicKey, error) {
	if len(verifyPoints) == 0 {
		return nil, errors.New("verifyPoints is nil")
	}

	// 将所有潜在节点的秘密值的验证点相加
	publicPoint, err := ecc.NewPoint(elliptic.P256(), verifyPoints[0].X, verifyPoints[0].Y)
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(verifyPoints); i++ {
		publicPoint, err = publicPoint.Add(verifyPoints[i])
		if err != nil {
			return nil, err
		}
	}

	// 将秘密值的验证点相加之和，转换成椭圆曲线公钥
	tssPublickey := new(ecdsa.PublicKey)
	tssPublickey.Curve = publicPoint.Curve
	tssPublickey.X = publicPoint.X
	tssPublickey.Y = publicPoint.Y

	return tssPublickey, nil
}

// -----  以下函数仅供测试使用，请勿调用 -----

// 测试用
func CalculatePublicKey(verifyPoints []*ecc.Point) (publicKeyPoint *ecc.Point, err error) {
	eddsaPublicKeyPoint, err := ecc.NewPoint(elliptic.P256(), verifyPoints[0].X, verifyPoints[0].Y)
	if err != nil {
		return nil, errors.New("eddsaPublicKeyPoint is not on the curve")
	}

	return eddsaPublicKeyPoint, nil
}

// 测试用
// 产生本地秘密的私钥碎片，可以把每个碎片理解为一个坐标点。 key: partner index，也就是x坐标, value: 实际数值，也就是y坐标
// minimumShareNumber可以理解为threshold
func SecretShareLocalKeyGenerateWithVerifyPoints(totalShareNumber, minimumShareNumber int, secret []byte) (shares map[int]*big.Int, points []*ecc.Point, err error) {
	// 1. calculate "partial" key share
	//	secret, err := walletRand.GenerateEntropy(120)
	//	if err != nil {
	//		return nil, nil, err
	//	}

	curve := elliptic.P256()

	// 2. compute the shares and the verify points
	shares, points, err = secret_share.ComplexSecretSplitWithVerifyPoints(totalShareNumber, minimumShareNumber, secret, curve)
	if err != nil {
		return nil, nil, err
	}

	return shares, points, nil
}

// 测试用
// 从收集的所有碎片中保存与自己相关的密钥部分，并保存在本地
func SecretShareLocalKeyGather(allPartnerShares []*PartnerShares, localIndex int) (shares map[int]*big.Int) {
	shares = make(map[int]*big.Int, len(allPartnerShares))

	//	log.Printf("allPartnerShares is: %v", allPartnerShares)
	for _, partnerShares := range allPartnerShares {
		//		log.Printf("partnerShares is: %v", partnerShares)
		//		log.Printf("partnerShares.PartnerInfo.Index is: %v", partnerShares.PartnerInfo.Index)
		//		log.Printf("partnerShares.Shares[localIndex] is: %v", partnerShares.Shares[localIndex])
		shares[partnerShares.PartnerInfo.Index] = partnerShares.Shares[localIndex]
	}

	return shares
}

// 测试用
// 从收集的所有碎片中保存与自己相关的密钥部分，并保存在本地
func LocalPrivateSharesGather(allPartnerShares []*PartnerShares, localIndex int) []*big.Int {
	var shares []*big.Int

	//	log.Printf("allPartnerShares is: %v", allPartnerShares)
	for _, partnerShares := range allPartnerShares {
		shares = append(shares, partnerShares.Shares[localIndex])
	}

	return shares
}

// 测试用
// 从本地收集的所有碎片计算 自己的xi，这是一个关键秘密信息
func CalcuateXi(shares map[int]*big.Int) *big.Int {
	xi := big.NewInt(0)

	for _, value := range shares {
		xi = new(big.Int).Add(xi, value)
	}

	xi = new(big.Int).Mod(xi, elliptic.P256().Params().N)

	return xi
}

// 测试用
// 从收集的所有碎片中保存验证点部分，并保存在本地
func SecretShareVerifyPointsGather(allPartnerShares []*PartnerShares, threshold int) (points []*ecc.Point, err error) {
	verifyPoints := make([]*ecc.Point, threshold)

	for _, partnerShares := range allPartnerShares {
		for i := 0; i < threshold; i++ {
			if verifyPoints[i] == nil {
				verifyPoints[i] = partnerShares.VerifyPoints[i]
				continue
			}
			verifyPoints[i], err = verifyPoints[i].Add(partnerShares.VerifyPoints[i])
			if err != nil {
				return nil, err
			}
		}
	}

	return verifyPoints, nil
}

// 测试用
func CalculatePublicKeys(verifyPoints []*ecc.Point, allPartnerShares []*PartnerShares, threshold int) (publicKeysPoints []*ecc.Point, err error) {
	publicKeysPoints = make([]*ecc.Point, len(allPartnerShares))

	for _, partnerShares := range allPartnerShares {
		bigXPoint := verifyPoints[0]
		for c := 1; c < threshold; c++ {
			z := partnerShares.PartnerInfo.IndentityKey
			bigXPoint, err = bigXPoint.Add(verifyPoints[c].ScalarMult(z))
			if err != nil {
				return nil, err
			}
		}
		publicKeysPoints[partnerShares.PartnerInfo.Index-1] = bigXPoint
	}

	return publicKeysPoints, nil
}

// -----  以上函数仅供测试使用，请勿调用 -----
