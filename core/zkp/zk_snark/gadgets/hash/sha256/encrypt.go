package sha256

import (
	"github.com/consensys/gnark/frontend"
)

// TODO: execution of a SHA256 run expressed as r1cs
func encryptBLS381(api frontend.API, h SHA256, message frontend.Variable, key frontend.Variable) frontend.Variable {
	res := message
	return res
}

// newSHA256BLS381 creates new SHA256 object
func newSHA256BLS381(seed string, api frontend.API) SHA256 {
	res := SHA256{}
	return res
}
