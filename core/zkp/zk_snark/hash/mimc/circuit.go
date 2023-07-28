package mimc

import (
	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/constraint"
	"github.com/consensys/gnark/examples/mimc"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/frontend/cs/r1cs"
)

// NewConstraintSystem return the compiled ConstraintSystem implementing a pre image check
func NewConstraintSystem() (constraint.ConstraintSystem, error) {
	// gnark already defined mimic circuit
	// mimc(preImage) == hash
	circuit := &mimc.Circuit{}

	// generate CompiledConstraintSystem
	return frontend.Compile(ecc.BLS12_381.ScalarField(), r1cs.NewBuilder, circuit)
}
