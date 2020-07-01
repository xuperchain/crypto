package key

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"

	"github.com/xuperchain/crypto/common/account"
	"github.com/xuperchain/crypto/gm/hash"

	accountUtil "github.com/xuperchain/crypto/gm/account"
	aesUtil "github.com/xuperchain/crypto/gm/aes"
)

type AccountInfo struct {
	Address           string
	PrivateKey        []byte
	PublicKey         []byte
	Mnemonic          []byte
	EncryptPrivateKey []byte
	EncryptPublicKey  []byte
	EncryptMnemonic   []byte
}

// 通过助记词来恢复并保存私钥
// 这里不应该再需要知道指定曲线了，也不需要知道版本号了，这个功能应该由助记词中的标记位来判断
func CreateAndSaveSecretKeyWithMnemonic(path string, language int, mnemonic string, password string) (*account.ECDSAInfo, error) {
	// 通过助记词来产生钱包账户
	ecdsaAccount, err := accountUtil.GenerateAccountByMnemonic(mnemonic, language)
	if err != nil {
		return nil, err
	}

	// 将私钥加密后保存到指定路径
	err = savePrivateKey(path, password, ecdsaAccount)
	if err != nil {
		return nil, err
	}

	// 返回的字段：助记词、随机byte数组、钱包地址
	ecdasaInfo := getECDSAInfoFromECDSAAccount(ecdsaAccount)

	return ecdasaInfo, nil
}

// 使用支付密码加密账户信息
func EncryptAccount(info *account.ECDSAAccount, password string) (*account.ECDSAAccountToCloud, error) {
	if info.JsonPrivateKey == "" {
		return nil, ErrParam
	}

	// 将aes对称加密的密钥扩展至32字节
	//	newPassword := hash.DoubleSha256([]byte(password))
	newPassword := hash.HashUsingSM3([]byte(password))

	// 加密私钥
	encryptedPrivateKey, err := aesUtil.Encrypt([]byte(info.JsonPrivateKey), newPassword)
	if err != nil {
		return nil, err
	}

	accountToClound := new(account.ECDSAAccountToCloud)
	accountToClound.JsonEncryptedPrivateKey = string(encryptedPrivateKey)
	accountToClound.Password = password
	accountToClound.Address = info.Address

	// 加密助记词
	if info.Mnemonic != "" {
		encryptedMnemonic, err := aesUtil.Encrypt([]byte(info.Mnemonic), newPassword)
		if err != nil {
			return nil, err
		}

		accountToClound.EncryptedMnemonic = string(encryptedMnemonic)
	}

	return accountToClound, nil
}

// 保存账户信息到文件,只需要保存address 和 privateKey
func SaveAccountFile(account *account.ECDSAAccountToCloud, path string) error {
	//如果path不是以/结尾的，自动拼上
	if strings.LastIndex(path, "/") != len([]rune(path))-1 {
		path = path + "/"
	}
	err := writeFileUsingFilename(path+"address", []byte(account.Address))
	if err != nil {
		return err
	}

	err = writeFileUsingFilename(path+"private.key", []byte(account.JsonEncryptedPrivateKey))
	if err != nil {
		return err
	}
	return nil
}

// 生成并保存私钥
func CreateAndSaveSecretKey(path string, language int, strength uint8, password string, cryptography uint8) (*account.ECDSAInfo, error) {
	//函数向指定的文件中写入数据。如果文件不存在将创建文件，否则会在写入数据之前清空文件。
	ecdsaAccount, err := accountUtil.CreateNewAccountWithMnemonic(language, strength, cryptography)
	if err != nil {
		return nil, err
	}

	// 将私钥加密后保存到指定路径
	err = savePrivateKey(path, password, ecdsaAccount)
	if err != nil {
		return nil, err
	}

	// 返回的字段：助记词、随机byte数组、钱包地址
	ecdasaInfo := getECDSAInfoFromECDSAAccount(ecdsaAccount)

	return ecdasaInfo, err
}

// 剔除掉ECDSAAccount需要隐藏的数据，返回的字段：助记词、随机byte数组、钱包地址
func getECDSAInfoFromECDSAAccount(ecdsaAccount *account.ECDSAAccount) *account.ECDSAInfo {
	ecdasaInfo := new(account.ECDSAInfo)
	ecdasaInfo.Mnemonic = ecdsaAccount.Mnemonic
	ecdasaInfo.EntropyByte = ecdsaAccount.EntropyByte
	ecdasaInfo.Address = ecdsaAccount.Address

	return ecdasaInfo
}

// 将私钥加密后保存到指定路径
func savePrivateKey(path string, password string, ecdsaAccount *account.ECDSAAccount) error {
	//如果path不是以/结尾的，自动拼上
	if strings.LastIndex(path, "/") != len([]rune(path))-1 {
		path = path + "/"
	}

	// 将aes对称加密的密钥扩展至32字节
	//	newPassword := hash.DoubleSha256([]byte(password))
	newPassword := hash.HashUsingSM3([]byte(password))

	// 加密密钥文件
	encryptContent, err := aesUtil.Encrypt([]byte(ecdsaAccount.JsonPrivateKey), newPassword)
	if err != nil {
		log.Printf("encrypt private key failed, the err is %v", err)
		return err
	}

	//	log.Printf("Export mnemonic file is successful, the path is %v", path+"mnemonic")
	err = writeFileUsingFilename(path+"private.key", encryptContent)
	if err != nil {
		log.Printf("Export private key file failed, the err is %v", err)
		return err
	}

	return nil
}

// 保存私钥
func writeFileUsingFilename(filename string, content []byte) error {
	//函数向filename指定的文件中写入数据(字节数组)。如果文件不存在将按给出的权限创建文件，否则在写入数据之前清空文件。
	contentStr := base64.StdEncoding.EncodeToString(content)
	err := ioutil.WriteFile(filename, []byte(contentStr), 0666)
	return err
}
