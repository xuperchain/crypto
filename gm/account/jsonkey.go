package account

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
)

// 通过这个数据结构来生成私钥的json
type ECDSAPrivateKey struct {
	Curvname string
	X, Y, D  *big.Int
}

// 通过这个数据结构来生成公钥的json
type ECDSAPublicKey struct {
	Curvname string
	X, Y     *big.Int
}

func getNewEcdsaPrivateKey(k *ecdsa.PrivateKey) *ECDSAPrivateKey {
	key := new(ECDSAPrivateKey)
	key.Curvname = k.Params().Name
	key.D = k.D
	key.X = k.X
	key.Y = k.Y

	return key
}

func getNewEcdsaPublicKey(k *ecdsa.PrivateKey) *ECDSAPublicKey {
	key := new(ECDSAPublicKey)
	key.Curvname = k.Params().Name
	key.X = k.X
	key.Y = k.Y

	return key
}

func getNewEcdsaPublicKeyFromPublicKey(k *ecdsa.PublicKey) *ECDSAPublicKey {
	key := new(ECDSAPublicKey)
	key.Curvname = k.Params().Name
	key.X = k.X
	key.Y = k.Y

	return key
}

// 获得私钥所对应的的json
func GetEcdsaPrivateKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	// 转换为自定义的数据结构
	key := getNewEcdsaPrivateKey(k)

	// 转换json
	data, err := json.Marshal(key)

	return string(data), err
}

// 获得公钥所对应的的json
func GetEcdsaPublicKeyJsonFormat(k *ecdsa.PrivateKey) (string, error) {
	// 转换为自定义的数据结构
	key := getNewEcdsaPublicKey(k)

	// 转换json
	data, err := json.Marshal(key)

	return string(data), err
}

// 获得公钥所对应的的json
func GetEcdsaPublicKeyJsonFormatFromPublicKey(k *ecdsa.PublicKey) (string, error) {
	// 转换为自定义的数据结构
	key := getNewEcdsaPublicKeyFromPublicKey(k)

	// 转换json
	data, err := json.Marshal(key)

	return string(data), err
}
