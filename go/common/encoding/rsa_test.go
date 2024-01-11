package encoding

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"testing"
)

func init() {
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048) // 生成私钥
}

func TestPrintPrivateKey(t *testing.T) {
	pubASN1 := x509.MarshalPKCS1PrivateKey(privateKey)
	pemData := string(pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: pubASN1,
	}))
	t.Log(pemData)

	LoadRsaPrivateKey(pemData)

	TestEncodeRsa(t)
}

func TestGetRsaPublicKey(t *testing.T) {
	pemData := GetRsaPublicKey()
	t.Log(pemData)
	rsaPubKey := PaserRsaPublicKey(pemData)

	// encrypt
	message := []byte("123")
	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPubKey, message)
	if err != nil {
		t.Fatal(err)
	}
	encrypt := base64.StdEncoding.EncodeToString(encryptedData)
	// decrypt
	data, err := DecodeRsa(encrypt)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)

}

func TestEncodeRsa(t *testing.T) {
	data, err := EncodeRsa([]byte("abc"))
	if err != nil {
		t.Fatal(err)
	}

	data, err = DecodeRsa(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(data)
}
