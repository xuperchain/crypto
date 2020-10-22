package ecc

import (
	"crypto/elliptic"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/xuperchain/crypto/gm/gmsm/sm2"
)

type Point struct {
	Curve elliptic.Curve
	X     *big.Int
	Y     *big.Int
}

type ECPoint struct {
	Curvname string
	X, Y     *big.Int
}

func NewPoint(curve elliptic.Curve, x, y *big.Int) (*Point, error) {
	if !curve.IsOnCurve(x, y) {
		return nil, errors.New("NewPoint: the given point is not on the elliptic curve")
	}
	return &Point{Curve: curve, X: x, Y: y}, nil
}

func newPoint(curve elliptic.Curve, x, y *big.Int) *Point {
	return &Point{Curve: curve, X: x, Y: y}
}

func (p *Point) ToString() (string, error) {
	// 转换为自定义的数据结构
	point := getECPoint(p)

	// 转换json
	data, err := json.Marshal(point)

	return string(data), err
}

func getECPoint(p *Point) *ECPoint {
	point := new(ECPoint)
	point.Curvname = p.Curve.Params().Name
	point.X = p.X
	point.Y = p.Y

	return point
}

func NewPointFromString(pointStr string) (*Point, error) {
	jsonContent := []byte(pointStr)
	ecPoint := new(ECPoint)
	err := json.Unmarshal(jsonContent, ecPoint)
	if err != nil {
		return nil, err //json有问题
	}

	curve := elliptic.P256()
	if ecPoint.Curvname != "P-256" && ecPoint.Curvname != "SM2-P-256" {
		err = fmt.Errorf("curve [%v] is not supported yet.", ecPoint.Curvname)
		return nil, err
	}
	if ecPoint.Curvname == "SM2-P-256" {
		curve = sm2.P256Sm2()
	}

	return NewPoint(curve, ecPoint.X, ecPoint.Y)
}

func (p *Point) Add(p1 *Point) (*Point, error) {
	x, y := p.Curve.Add(p.X, p.Y, p1.X, p1.Y)

	return NewPoint(p.Curve, x, y)
}

func (p *Point) ScalarMult(k *big.Int) *Point {
	x, y := p.Curve.ScalarMult(p.X, p.Y, k.Bytes())
	newP := newPoint(p.Curve, x, y)

	return newP
}

func ScalarBaseMult(curve elliptic.Curve, k *big.Int) *Point {
	x, y := curve.ScalarBaseMult(k.Bytes())
	p := newPoint(curve, x, y)

	return p
}

func (p *Point) Equals(p1 *Point) bool {
	if p == nil || p1 == nil {
		return false
	}

	if p.X.Cmp(p1.X) != 0 || p.Y.Cmp(p1.Y) != 0 {
		return false
	}

	return true
}
