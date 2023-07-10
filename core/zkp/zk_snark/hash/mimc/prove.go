package mimc

import (
	"github.com/consensys/gnark-crypto/ecc"
	groth16_bls381 "github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/examples/mimc"
	"github.com/consensys/gnark/frontend"

	"github.com/xuperchain/crypto/core/hash"
)

// Prove generate a zkp proof using ProvingKey
func Prove(ccs constraint.ConstraintSystem, pk groth16_bls381.ProvingKey, secret []byte) (groth16_bls381.Proof, error) {
	hashResult := hash.HashUsingDefaultMiMC(secret)
	assignment := &mimc.Circuit{
		PreImage: secret,
		Hash:     hashResult,
	}
	witness, err := frontend.NewWitness(assignment, ecc.BLS12_381.ScalarField())
	if err != nil {
		return nil, err
	}

	proof, err := groth16_bls381.Prove(ccs, pk, witness)
	if err != nil {
		return nil, err
	}

	return proof, nil
}
