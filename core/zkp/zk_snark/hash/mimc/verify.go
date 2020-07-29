package mimc

import (
	"github.com/consensys/gnark/backend"
	"log"

	//	groth16_bls377 "github.com/consensys/gnark/backend/bls377/groth16"
	//	groth16_bls381 "github.com/consensys/gnark/backend/bls381/groth16"
	groth16_bn256 "github.com/consensys/gnark/backend/bn256/groth16"
)

func Verify(proof *groth16_bn256.Proof, vk *groth16_bn256.VerifyingKey, hashResult []byte) (bool, error) {
	//	var assignment backend.Assignments
	//	assignment.Assign(backend.Visibility("public"), "hash", hashResult)

	assignment := backend.NewAssignment()
	assignment.Assign(backend.Public, "hash", hashResult)

	isValid, err := groth16_bn256.Verify(proof, vk, assignment)
	if err != nil {
		log.Printf("Error during verify is: %v", err)
		return false, err
	}

	return isValid, nil
}
