package encoding

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

var privateKey *rsa.PrivateKey // 服务器私钥

// LoadRsaPrivateKey 加载rsa私钥
func LoadRsaPrivateKey(pemData []byte) {
	block, _ := pem.Decode(pemData)
	pubKey, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	privateKey = pubKey
}

// PaserRsaPublicKey 解析rsa公钥
func PaserRsaPublicKey(pemData string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(pemData))
	pubKey, _ := x509.ParsePKIXPublicKey(block.Bytes)
	return pubKey.(*rsa.PublicKey)
}

// GetRsaPublicKey 获取公钥
func GetRsaPublicKey() string {
	pubASN1, _ := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	return string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubASN1,
	}))
}

// EncodeRsa 加密
func EncodeRsa(data []byte) (string, error) {
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, &privateKey.PublicKey, data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptedData), nil
}

// DecodeRsa 解密
func DecodeRsa(data string) (string, error) {
	encrypt, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	plain, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encrypt)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}
