package account

import (
	"encoding/json"
	//"log"
	"encoding/hex"
	"math/big"

	"github.com/xuperchain/crypto/common/utils"
	"github.com/xuperchain/crypto/gm/secret_share/complex_secret_share"
)

//ShareGroup 密码学切分后的密钥片段簇
type ShareGroup []*big.Int

//PrivateKeyShare 私钥片段
type PrivateKeyShare struct {
	Index         int
	ComplexShares ShareGroup
}

//SplitPrivateKey 私钥分割
func SplitPrivateKey(jsonPrivKey string, totalShareNumber, minimumShareNumber int) ([]string, error) {
	privateKeyBytes := []byte(jsonPrivKey)
	//log.Printf("[SharePrivateKey]privateKeyBytes is %v", privateKeyBytes)

	privateKeyLength := len(privateKeyBytes)
	privateKeyByteGroup := [][]byte{}

	//按照32长度切分密钥，保证每部分都小于密钥分存算法支持的大数上限
	for i := 0; i < privateKeyLength; i += 32 {
		end := i + 32
		if end > privateKeyLength {
			end = privateKeyLength
		}
		privateKeyByteGroup = append(privateKeyByteGroup, privateKeyBytes[i:end])
	}
	//log.Printf("[SharePrivateKey]privateKeyByteGroup is %v", privateKeyByteGroup)

	//基于密码学对每段密钥进行切分
	complexShareGroup := map[int]ShareGroup{}
	for _, complexSecretMsg := range privateKeyByteGroup {
		complexShares, err := complex_secret_share.ComplexSecretSplit(totalShareNumber, minimumShareNumber, complexSecretMsg)
		if nil != err {
			return []string{}, err
		}

		for key, share := range complexShares {
			complexShareGroup[key] = append(complexShareGroup[key], share)
		}
	}

	//构建返回值结构
	result := []string{}
	for key, cs := range complexShareGroup {
		encoRet, err := complexShareEncode(cs, key)
		if nil != err {
			return []string{}, err
		}
		result = append(result, encoRet)
	}

	return result, nil
}

//RetrievePrivateKeyByShares 私钥恢复
func RetrievePrivateKeyByShares(strEncodeComplexShareGroup []string) (string, error) {
	//对输入进行解码，获取所有的密钥片段簇
	complexShareGroup := map[int]ShareGroup{}
	for _, ecs := range strEncodeComplexShareGroup {
		shareGroup, index, err := complexShareDecode(ecs)
		if nil != err {
			return "", err
		}
		complexShareGroup[index] = shareGroup
	}

	complexShareFragments := map[int](map[int]*big.Int){}
	for key, shares := range complexShareGroup {
		for i, v := range shares {
			if _, created := complexShareFragments[i+1]; created {
				complexShareFragments[i+1][key] = v
			} else {
				complexShareFragments[i+1] = map[int]*big.Int{key: v}
			}
		}
	}
	//log.Printf("[SharePrivateKey]complexShareFragments is: %v", complexShareFragments)

	secretFragments := [][]byte{}
	for key := 1; true; key++ {
		complexShares, exits := complexShareFragments[key]
		if !exits {
			break
		}

		retrieveComplexShares := map[int]*big.Int{}

		for k, v := range complexShares {
			retrieveComplexShares[k] = v
		}

		secretBytes, err := complex_secret_share.ComplexSecretRetrieve(retrieveComplexShares)
		if nil != err {
			return "", err
		}

		secretFragments = append(secretFragments, secretBytes)
	}
	//log.Printf("[SharePrivateKey]secretFragments is: %v", secretFragments)

	secretBytes := utils.BytesCombine(secretFragments...)
	//log.Printf("[SharePrivateKey] ComplexSecretRetrieve bytes is: %v", secretBytes)
	privateKey := string(secretBytes)
	//log.Printf("[SharePrivateKey] ComplexSecretRetrieve privatekey is: %s", privateKey)
	return privateKey, nil
}

//complexShareEncode 对一簇密钥片段进行编码
func complexShareEncode(shares ShareGroup, index int) (string, error) {
	pks := PrivateKeyShare{
		Index:         index,
		ComplexShares: shares,
	}
	jr, err := json.Marshal(pks)
	if nil != err {
		return "", err
	}
	return hex.EncodeToString(jr), nil
}

//complexShareDecode 解密获得一簇密钥片段
func complexShareDecode(hexCode string) (ShareGroup, int, error) {
	codeBytes, err := hex.DecodeString(hexCode)
	if nil != err {
		return nil, 0, err
	}
	pks := PrivateKeyShare{}
	err = json.Unmarshal(codeBytes, &pks)
	if nil != err {
		return nil, 0, err
	}
	return pks.ComplexShares, pks.Index, nil
}
