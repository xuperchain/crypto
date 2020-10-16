package rand

import (
	"errors"
	"fmt"
)

var (
	// 熵的长度不在 [128, 256]以内或者长度不是32的倍数
	ErrInvalidEntropyLength = errors.New("Entropy length must within [128, 256] and be multiples of 32")

	// 助记词的强度暂未被支持
	// Strength required for generating Mnemonic not supported yet.
	ErrStrengthNotSupported = fmt.Errorf("This strength has not been supported yet.")
)
