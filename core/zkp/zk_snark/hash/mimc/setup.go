package mimc

import (
	groth16_bls381 "github.com/consensys/gnark/backend/groth16"
	"github.com/xuperchain/crypto/common/zkp"
)

// Setup generate CompiledConstraintSystem, ProvingKey and VerifyingKey
func Setup() (*zkp.ZkpInfo, error) {
	r1cs, err := NewConstraintSystem()
	if err != nil {
		return nil, err
	}

	pk, vk, err := groth16_bls381.Setup(r1cs)
	if err != nil {
		return nil, err
	}

	zkpInfo := &zkp.ZkpInfo{
		R1CS:         r1cs,
		ProvingKey:   pk,
		VerifyingKey: vk,
	}

	return zkpInfo, nil
}
