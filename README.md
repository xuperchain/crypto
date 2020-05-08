# XuperChain Crypto

This project is the crypto library for XuperChain.

本项目是超级链相关的密码库模块。目前开源部分已经支持了Xuper统一超级签名算法、Schnorr签名，环签名、多重签名等多种签名算法。也支持了秘密分享、分层确定性加解密等多种密码学能力。

NIST系列算法请引用：
"github.com/xuperchain/crypto/client/service/xchain"

详细函数说明请参考该类的函数注释

使用例子：
```
import (
	"log"

	"github.com/xuperchain/crypto/client/service/xchain"
	"github.com/xuperchain/crypto/core/account"
	"github.com/xuperchain/crypto/core/hdwallet/rand"
)

xcc := new(xchain.XchainCryptoClient)

ecdsaAccount, err := xcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
if err != nil {
	log.Printf("CreateNewAccountWithMnemonic failed and err is: %v", err)
	return
}

log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)
```

------

国密系列算法请引用：
"github.com/xuperchain/crypto/client/service/gm"

详细函数说明请参考该类的函数注释

使用例子：
```
import (
	"log"

	"github.com/xuperchain/crypto/client/service/gm"
	"github.com/xuperchain/crypto/gm/account"
	"github.com/xuperchain/crypto/gm/hdwallet/rand"
)

gcc := new(gm.GmCryptoClient)

ecdsaAccount, err := gcc.CreateNewAccountWithMnemonic(rand.SimplifiedChinese, account.StrengthHard)
if err != nil {
	log.Printf("CreateNewAccountWithMnemonic failed and err is: %v", err)
	return
}
log.Printf("mnemonic is %v, jsonPrivateKey is %v, jsonPublicKey is %v and address is %v", ecdsaAccount.Mnemonic, ecdsaAccount.JsonPrivateKey, ecdsaAccount.JsonPublicKey, ecdsaAccount.Address)
```