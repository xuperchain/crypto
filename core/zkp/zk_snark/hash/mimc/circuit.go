package mimc

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/examples/mimc"
	"github.com/consensys/gnark/frontend"
)

// NewCircuit return the circuit implementing a pre image check
func NewCircuit() (frontend.CompiledConstraintSystem, error) {
	// gnark already defined mimic circuit
	// mimc(preImage) == hash
	circuit := &mimc.Circuit{}

	// generate CompiledConstraintSystem
	ccs, err := frontend.Compile(ecc.BLS12_381, backend.GROTH16, circuit)
	if err != nil {
		return nil, err
	}

	return ccs, nil
}
