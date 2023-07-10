package mimc

import (
	"github.com/consensys/gnark-crypto/ecc"
	groth16_bls381 "github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/examples/mimc"
	"github.com/consensys/gnark/frontend"
)

// Verify verify a zkp proof using VerifyingKey
func Verify(proof groth16_bls381.Proof, vk groth16_bls381.VerifyingKey, hashResult []byte) (bool, error) {
	assignment := &mimc.Circuit{
		Hash: hashResult,
	}
	witness, err := frontend.NewWitness(assignment, ecc.BLS12_381.ScalarField())
	if err != nil {
		return false, err
	}

	if err := groth16_bls381.Verify(proof, vk, witness); err != nil {
		return false, err
	}
	return true, nil
}
