package hash

import (
	"github.com/xuperchain/crypto/gm/gmsm/sm3"
)

func HashUsingSM3(data []byte) []byte {
	var h sm3.SM3

	h.Reset()
	h.Write(data)
	out := h.Sum(nil)

	return out
}
