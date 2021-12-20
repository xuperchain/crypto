package sha256

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/frontend"

	"github.com/xuperchain/crypto/core/zkp/zk_snark/gadgets/hash/sha256"
)

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
func (circuit *SHA256Circuit) Define(curveID ecc.ID, api frontend.API) error {
	// hash function
	sha256, _ := sha256.NewSHA256("seed", curveID, api)

	// specify constraints
	// SHA256(preImage) == hash
	sha256.Write(circuit.PreImage)
	api.AssertIsEqual(circuit.Hash, sha256.Sum())
	return nil
}

// NewCircuit return the circuit implementing a pre image check
func NewCircuit() (frontend.CompiledConstraintSystem, error) {
	circuit := &SHA256Circuit{}

	// generate CompiledConstraintSystem
	ccs, err := frontend.Compile(ecc.BLS12_381, backend.GROTH16, circuit, nil)
	if err != nil {
		return nil, err
	}

	return ccs, nil
}
