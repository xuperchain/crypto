package sha256

import (
	"math/big"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/frontend"
)

// SHA256 contains the params of the SHA256 hash func and the curves on which it is implemented
type SHA256 struct {
	params []big.Int           // slice containing constants for the encryption rounds
	id     ecc.ID              // id needed to know which encryption function to use
	h      frontend.Variable   // current vector in the Miyaguchi–Preneel scheme
	data   []frontend.Variable // state storage. data is updated when Write() is called. Sum sums the data.
	api    frontend.API        // underlying constraint system
}

// NewSHA256 returns a SHA256 instance, than can be used in a gnark circuit
func NewSHA256(seed string, id ecc.ID, api frontend.API) (SHA256, error) {
	// only support bls381 currently
	return newSHA256BLS381(seed, api), nil
}

// Write adds more data to the running hash.
func (h *SHA256) Write(data ...frontend.Variable) {
	h.data = append(h.data, data...)
}

// Reset resets the Hash to its initial state.
func (h *SHA256) Reset() {
	h.data = nil
	h.h = h.api.Constant(0)
}

// Hash hash (in r1cs form) using Miyaguchi–Preneel:
// https://en.wikipedia.org/wiki/One-way_compression_function
// The XOR operation is replaced by field addition.
// See github.com/consensys/gnark-crypto for reference implementation.
func (h *SHA256) Sum() frontend.Variable {
	for _, stream := range h.data {
		h.h = encryptBLS381(h.api, *h, stream, h.h)
		h.h = h.api.Add(h.h, stream)
	}

	h.data = nil // flush the data already hashed

	return h.h
}
