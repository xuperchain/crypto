package complex_secret_share

import (
	"errors"
	"log"
	"math/big"

	polynomial "github.com/xuperchain/crypto/gm/secret_share/big_polynomial"
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
func ComplexSecretSplit(totalShareNumber, minimumShareNumber int, secret []byte) (shares map[int]*big.Int, err error) {
	// Check the parameters
	if totalShareNumber < 2 {
		return nil, InvaildTotalShareNumberError
	}

	if minimumShareNumber > totalShareNumber {
		return nil, InvaildShareNumberError
	}

	poly, err := polynomial.RandomGenerate(minimumShareNumber-1, secret)
	if err != nil {
		return nil, err
	}

	log.Printf("The asc order coefficients of the polynomial is: %v", poly)

	// Evaluate the polynomial for several times, in order to get all the shares.
	shares = make(map[int]*big.Int, totalShareNumber)
	//	for x := 1; x <= totalShareNumber; x++ {
	//		shares[x] = polynomial.Evaluate(poly, big.NewInt(int64(x)))
	//	}
	for x := 1; x <= totalShareNumber; x++ {
		shares[x] = polynomial.Evaluate(poly, big.NewInt(int64(x)))
	}
	log.Printf("shares is: %v", shares)

	return shares, nil
}
