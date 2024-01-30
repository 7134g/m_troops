package encoding

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
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
	fmt.Println(pemData)

	LoadRsaPrivateKey([]byte(pemData))

	TestEncodeRsa(t)
}

func TestLoadRsaPrivateKey(t *testing.T) {
	pemData := `-----BEGIN PRIVATE KEY-----
MIIEowIBAAKCAQEAuIz6zqLXAJhlWoVZQydPOjeaHssE0hVfCU2SS+5JFurNA8oV
x34mL3gxFz+P/1hEx2NsU4TcbkA7H6IaN8x1iY/x07vjvYwjdxl8aH2y2TnleFfA
pipR/T7v4GMqpzePk5NVi+qlbBTdvPuD7V6DkawnrpPkxfUkJOcddVr0R5jy1+Kp
g9Ix0PtgKD0U/8NEYhJ+gwWvUEbg3VjAJdYrlZObA2fS78CLGpuSDLbU+y4+fJqb
2UMoGDPRibO8SZoZWqUxKVdZMLUasYg5Eo9pzRzQHRiyTMsr3EzbYLmLl6Lm2mZQ
kj4xg1YlHL8D/gl7R1qePpNR9VtI1/HvK1NLYQIDAQABAoIBAFqldHOWbEBwmifS
I7vmTPXtZZgGZAHEWX+6SEdfbSsCfVyFIBlyjtL2GDaxE8t50Z6V6RlEwvisN94b
wwDxPjIpL8CuIvhxIuJW9FJmmzgzRdDMWWXEl6zqAuyWCNdFZRI5ZeVphYMG5Lr0
VyJ6L+mXQ45uOgo5lF1D36ZK3H1lfyWTvYhdCxPL1/wgyU4hryTfPwzSb1Ev8p8m
zjukUTyEepSV0aAAfVzSbr8WKf49B82/dqBm4kM7grFZ+/+zmCy/SBclOTvL4JoO
FHpQNc/jOYiB4pwmjaLkhU5B6hz3AVJqf85FdhaNPMIKVrQwEuvMSHYpJWOzQ8mX
QkcogAECgYEA1h8R5dBBT3GfOUBEmSeweuI5Ug6Bwab52sQKq/NyJmydX7LuG5uD
4sP6Rk5t/FACQ9qmlgxPOPdwtWt7D++ltXlIdKHC+1VnsQIkvFOD5a9piZqy0w7Q
r44mjr3AocwWWte6ON+r0skTZvU8iBaJNPuMZmMjZbQ1YkFjoDHRzgECgYEA3KVS
+XYeQeDzfsC1pzH+3PqzJOIwM0yDgHu/wsXQodreKGj7lVcdbJIeVToR/tKoO6Hp
TKEOsqdzxgOvfjIphAXSdnZT1bJqkIbrYVJroeP0cfMX7QgtNcjP4gP0aLbpHXCI
irhuSxmaoGKGydpvuVeQ6oVT3nk50Y3c9mi+PWECgYA9TVnnbM81+na2gmLhYk6R
b/EvP/4APljsPBI+Fo3I2HHZ8zVebBC/PJKLzGqKwTFU0eW9sbqAub7oAeSStG7B
3P5Uffd/03zDXbS8wkBR4v2ZKtQlvukaSd1aIpNi/zYrPfYP0GG2EGFgqbdx0tnn
dFlG+v0oYgaiocvvEjRaAQKBgCTTSufruN0R6FHNZAQFqh3DkcakIZtON9xnyvd9
AHcjClUAQI0KPxTxxjI/QOWgzwc03LU3ZDaZEA+Kae3L/XXVauzujstpvbNlcT+K
+//HBfNGuUWMSc9iNp/oPRCFBp8tOvy8D1xlZ5NBHnHuDRuH693YZskIvoek/634
iVfBAoGBAJJZ9oxbRqZJYRQ0DC/SSOWDdvLfT8mqvhHk1w6O+tw/xF1FikuSa0EW
7wUipOiR8mnIYt1UrxDL+OOSZCnGs6G2DQ0ujgNxOAnytE6l1t5NW9/weWg3hRvR
7m7ECqs0J6/0kuQuds0ETBU/7aXqZstTuX5zRUAB5s/NI3CPNm44
-----END PRIVATE KEY-----
`

	data := `
{
    "region":"86",
    "phone":"19900000012",
    "password":"000000"
}
`

	LoadRsaPrivateKey([]byte(pemData))
	encrypt, err := EncodeRsa([]byte(data))
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(encrypt)
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
