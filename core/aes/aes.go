package aes

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/xuperchain/crypto/common/utils"
)

func Encrypt(originalData, key []byte) ([]byte, error) {
	// The key argument should be the AES key,
	// either 16, 24, or 32 bytes to select
	// AES-128, AES-192, or AES-256.
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	originalData = utils.BytesPKCS5Padding(originalData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cipherInfo := make([]byte, len(originalData))

	blockMode.CryptBlocks(cipherInfo, originalData)

	return cipherInfo, nil
}

func Decrypt(cipherInfo, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originalData := make([]byte, len(cipherInfo))

	blockMode.CryptBlocks(originalData, cipherInfo)

	return utils.BytesPKCS5UnPadding(originalData)
}
