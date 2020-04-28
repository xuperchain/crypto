package rand

import (
	"crypto/rand"
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"
)

// 定义不同int类型对应的key length
const (
	// int8 类型
	KeyLengthInt8 = 8

	// int16 类型
	KeyLengthInt16 = 16

	// int32 类型
	KeyLengthInt32 = 32

	// int64 类型
	KeyLengthInt64 = 64
)

const (
	// 安全强度低
	KeyStrengthEasy = iota

	// 安全强度中
	KeyStrengthMiddle

	// 安全强度高
	KeyStrengthHard
)

// 底层调用跟操作系统相关的函数（读取系统熵）来产生一些伪随机数，
// 对外建议管这个返回值叫做“熵”
func generateEntropy(bitSize int) ([]byte, error) {
	err := validateEntropyBitSize(bitSize)
	if err != nil {
		return nil, err
	}

	entropy := make([]byte, bitSize/8)
	_, err = rand.Read(entropy)
	return entropy, err
}

// 生成一个指定长度的随机数种子
func generateSeedWithRandomPassword(randomPassword []byte, keyLen int) []byte {
	salt := "jingbo is handsome."
	seed := pbkdf2.Key(randomPassword, []byte(salt), 2048, keyLen, sha512.New)

	return seed
}

func GenerateSeedWithStrengthAndKeyLen(strength int, keyLength int) ([]byte, error) {
	var entropyBitLength = 0
	//根据强度来判断随机数长度
	switch strength {
	case KeyStrengthEasy: // 弱
		entropyBitLength = 128
	case KeyStrengthMiddle: // 中
		entropyBitLength = 192
	case KeyStrengthHard: // 高
		entropyBitLength = 256
	default: // 不支持的语言类型
		entropyBitLength = 0
	}

	// 判断强度是否合法
	if entropyBitLength == 0 {
		return nil, ErrStrengthNotSupported
	}

	// 产生随机熵
	entropyByte, err := generateEntropy(entropyBitLength)
	if err != nil {
		return nil, err
	}

	return generateSeedWithRandomPassword(entropyByte, keyLength), nil
}
