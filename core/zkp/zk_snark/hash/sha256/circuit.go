package sha256

import (
	"github.com/consensys/gnark/frontend"
	//	"github.com/consensys/gurvy"
	"github.com/xuperchain/crypto/core/zkp/zk_snark/gadgets/hash/sha256"
)

// New return the circuit implementing
// a pre image check
func NewCircuit() *frontend.CS {
	// create root constraint system
	circuit := frontend.New()

	// declare secret and public inputs
	preInput := circuit.SECRET_INPUT("secret_msg")
	hash := circuit.PUBLIC_INPUT("hash")

	// hash function
	//	mimc, _ := mimc.NewMiMCGadget("seed", gurvy.BN256)
	sha256Hash := sha256.NewSHA256Gadget()

	// specify constraints
	circuit.MUSTBE_EQ(hash, sha256Hash.Hash(&circuit, preInput))

	return &circuit
}
