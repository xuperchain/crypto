package mimc

import (
	"log"

	"github.com/xuperchain/crypto/core/hash"

	"github.com/consensys/gnark/backend"
	//	"github.com/consensys/gnark/frontend"

	//	groth16_bls377 "github.com/consensys/gnark/backend/bls377/groth16"
	//	groth16_bls381 "github.com/consensys/gnark/backend/bls381/groth16"
	groth16_bn256 "github.com/consensys/gnark/backend/bn256/groth16"

	backend_bn256 "github.com/consensys/gnark/backend/bn256"
)

func Prove(r1cs *backend_bn256.R1CS, pk *groth16_bn256.ProvingKey, secret []byte) (*groth16_bn256.Proof, error) {
	//	assignment := backend.NewAssignment()
	//	var assignment backend.Assignments

	assignment := backend.NewAssignment()
	assignment.Assign(backend.Secret, "secret_msg", secret)

	//	assignment.Assign(backend.Visibility("secret"), "secret_msg", secret)

	hashResult := hash.HashUsingDefaultMiMC(secret)
	//	assignment.Assign(backend.Visibility("public"), "hash", hashResult)
	assignment.Assign(backend.Public, "hash", hashResult)
	log.Printf("hashResult of secret[%v] is: [%v]", secret, hashResult)

	proof, err := groth16_bn256.Prove(r1cs, pk, assignment)
	if err != nil {
		log.Printf("Error during proof generation is: %v", err)
		return nil, err
	}

	return proof, nil
}
