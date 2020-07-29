package mimc

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/gadgets/hash/mimc"
	"github.com/consensys/gurvy"
)

// New return the circuit implementing
// a pre image check
func NewCircuit() *frontend.CS {
	// create root constraint system
	circuit := frontend.New()

	// declare secret and public inputs
	preImage := circuit.SECRET_INPUT("secret_msg")
	hash := circuit.PUBLIC_INPUT("hash")

	// hash function
	mimc, _ := mimc.NewMiMCGadget("seed", gurvy.BN256)

	// specify constraints
	// mimc(preImage) == hash
	circuit.MUSTBE_EQ(hash, mimc.Hash(&circuit, preImage))

	//	r1cs := circuit.ToR1CS()

	return &circuit
}
