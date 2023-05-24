package encoding

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"log"
)

func AESEncrypt(origData, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	origData = PKCS7Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])

	crypt := make([]byte, len(origData))

	blockMode.CryptBlocks(crypt, origData)

	return crypt
}

func AESDecrypt(crypt, key []byte) []byte {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypt))
	blockMode.CryptBlocks(origData, crypt)
	origData = PKCS7UnPadding(origData)
	return origData
}

func PKCS7Padding(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(origData, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:length-unPadding]
}
