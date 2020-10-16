/*
Copyright © 2020 ConsenSys

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sha256

import (
	//	"math/big"

	"github.com/consensys/gnark/frontend"
	//	"github.com/consensys/gnark/gadgets"
	//	"github.com/consensys/gurvy"
)

// SHA256Gadget contains the params of the Mimc gadget and the curves on which it is implemented
type SHA256Gadget struct {
	//	Params []big.Int
	//	id     gurvy.ID
}

// NewMiMCGadget returns a MiMC gadget, than can be used in a circuit
func NewSHA256Gadget() SHA256Gadget {
	return SHA256Gadget{}
}

// Hash hash (in r1cs form) using Miyaguchi–Preneel:
// https://en.wikipedia.org/wiki/One-way_compression_function
// The XOR operation is replaced by field addition
func (h MiMCGadget) Hash(circuit *frontend.CS, data ...*frontend.Constraint) *frontend.Constraint {
	digest := circuit.ALLOCATE(0)

	for _, stream := range data {
		digest = encryptFuncs(circuit, h, stream, digest)
		digest = circuit.ADD(digest, stream)
	}

	return digest

}
