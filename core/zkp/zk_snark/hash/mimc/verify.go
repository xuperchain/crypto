package mimc

import (
	groth16_bls381 "github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/examples/mimc"
	"github.com/consensys/gnark/frontend"
)

// Verify verify a zkp proof using VerifyingKey
func Verify(proof groth16_bls381.Proof, vk groth16_bls381.VerifyingKey, hashResult []byte) (bool, error) {
	assignment := &mimc.Circuit{
		Hash: frontend.Value(hashResult),
	}

	if err := groth16_bls381.Verify(proof, vk, assignment); err != nil {
		return false, err
	}
	return true, nil
}
