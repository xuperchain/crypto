package zkp

import (
	//	groth16_bls377 "github.com/consensys/gnark/backend/bls377/groth16"
	//	groth16_bls381 "github.com/consensys/gnark/backend/bls381/groth16"
	groth16_bn256 "github.com/consensys/gnark/backend/bn256/groth16"

	backend_bn256 "github.com/consensys/gnark/backend/bn256"
)

// R1CS、ProvingKey、VerifyingKey
type ZkpInfo struct {
	R1CS         *backend_bn256.R1CS
	ProvingKey   *groth16_bn256.ProvingKey
	VerifyingKey *groth16_bn256.VerifyingKey
}
