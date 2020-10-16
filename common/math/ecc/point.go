package ecc

import (
	//	"bytes"
	"crypto/elliptic"
	//	"encoding/binary"
	//	"encoding/json"
	"errors"
	//	"fmt"
	"math/big"
)

type Point struct {
	Curve elliptic.Curve
	X     *big.Int
	Y     *big.Int
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
