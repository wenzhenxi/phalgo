package phalgo

//	PhalGo-des
//	des加解密和3des加解密
//	喵了个咪 <wenzhenxi@vip.qq.com> 2016/5/11
//  依赖情况:无依赖

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

type Des struct {

}


func (this *Des)DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = this.PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func (this *Des)DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = this.PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func (this *Des)ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func (this *Des)ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimRightFunc(origData, func(r rune) bool {
		return r == rune(0)
	})
}

// 3DES加密
func (this *Des)TripleDesEncrypt(origData []byte, key string, iv string) ([]byte, error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	origData = this.PKCS5Padding(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	//blockMode := cipher.NewCBCEncrypter(block, key[:8])
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

// 3DES解密
func (this *Des)TripleDesDecrypt(crypted []byte, key string, iv string) ([]byte, error) {
	block, err := des.NewTripleDESCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	//blockMode := cipher.NewCBCDecrypter(block, key[:8])
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = this.PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func (this *Des)PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (this *Des)PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length - 1])
	return origData[:(length - unpadding)]
}