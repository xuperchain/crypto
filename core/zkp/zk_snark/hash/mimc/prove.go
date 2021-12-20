package mimc

import (
	groth16_bls381 "github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/examples/mimc"
	"github.com/consensys/gnark/frontend"

	"github.com/xuperchain/crypto/core/hash"
)

// Prove generate a zkp proof using ProvingKey
func Prove(ccs frontend.CompiledConstraintSystem, pk groth16_bls381.ProvingKey, secret []byte) (groth16_bls381.Proof, error) {
	hashResult := hash.HashUsingDefaultMiMC(secret)
	assignment := &mimc.Circuit{
		PreImage: frontend.Value(secret),
		Hash:     frontend.Value(hashResult),
	}

	proof, err := groth16_bls381.Prove(ccs, pk, assignment)
	if err != nil {
		return nil, err
	}

	return proof, nil
}
