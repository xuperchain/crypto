package mimc

import (
	"errors"

	"github.com/xuperchain/crypto/common/zkp"

	//	"github.com/consensys/gnark/frontend"

	//	groth16_bls377 "github.com/consensys/gnark/backend/bls377/groth16"
	//	groth16_bls381 "github.com/consensys/gnark/backend/bls381/groth16"
	groth16_bn256 "github.com/consensys/gnark/backend/bn256/groth16"

	backend_bn256 "github.com/consensys/gnark/backend/bn256"
)

var (
	InvalidInputParamsError = errors.New("Invalid input params")
)

// Todo: 优化返回值的数量，新增一个struct，超过2个返回值不友好
func Setup() *zkp.ZkpInfo {
	mimcCircuit := NewCircuit()

	//circuit to rank-1 constraint system
	r1cs := backend_bn256.New(mimcCircuit)

	return setup(&r1cs)
}

func setup(r1cs *backend_bn256.R1CS) *zkp.ZkpInfo {
	//	var pk groth16.ProvingKey
	//	var vk groth16.VerifyingKey

	var pk groth16_bn256.ProvingKey
	var vk groth16_bn256.VerifyingKey

	// 1: setup, needs to be run only once per circuit
	groth16_bn256.Setup(r1cs, &pk, &vk)

	zkpInfo := &zkp.ZkpInfo{
		R1CS:         r1cs,
		ProvingKey:   &pk,
		VerifyingKey: &vk,
	}

	return zkpInfo

	//	// 2: prove
	//	proof, err := groth16.Prove(r1cs, &pk, input)
	//
	//	// 3: verify
	//	err := groth16.Verify(proof, &vk, publicInput)

}
