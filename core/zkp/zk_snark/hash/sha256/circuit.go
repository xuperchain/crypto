package sha256

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"

	"github.com/xuperchain/crypto/core/zkp/zk_snark/gadgets/hash/sha256"
)

const useless ecc.ID = 0

// Circuit defines a pre-image knowledge proof
// SHA256(secret preImage) = public hash
type SHA256Circuit struct {
	// struct tag on a variable is optional
	// default uses variable name and secret visibility.
	PreImage frontend.Variable
	Hash     frontend.Variable `gnark:",public"`
}

// Define declares the circuit's constraints
// Hash = SHA256(PreImage)
func (circuit *SHA256Circuit) Define(api frontend.API) error {
	// hash function
	sha256, _ := sha256.NewSHA256("seed", useless, api)

	// specify constraints
	// SHA256(preImage) == hash
	sha256.Write(circuit.PreImage)
	api.AssertIsEqual(circuit.Hash, sha256.Sum())
	return nil
}

// NewConstraintSystem return the compiled ConstraintSystem implementing a pre image check
func NewConstraintSystem() (constraint.ConstraintSystem, error) {
	circuit := &SHA256Circuit{}

	// generate CompiledConstraintSystem
	return frontend.Compile(ecc.BLS12_381.ScalarField(), r1cs.NewBuilder, circuit)
}
