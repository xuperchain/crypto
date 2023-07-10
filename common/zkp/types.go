package zkp

import (
	groth16_bls381 "github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/constraint"
)

// ZkpInfo includes CompiledConstraintSystem、ProvingKey、VerifyingKey
type ZkpInfo struct {
	R1CS         constraint.ConstraintSystem
	ProvingKey   groth16_bls381.ProvingKey
	VerifyingKey groth16_bls381.VerifyingKey
}
