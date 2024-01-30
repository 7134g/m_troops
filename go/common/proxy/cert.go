package proxy

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"github.com/google/martian/mitm"
	"os"
	"time"
)

var (
	domain        = "proxy"                   // 域名
	organization  = "location"                // 组织者名称
	validDuration = time.Hour * 24 * 365 * 20 // 证书过期时间
)

var (
	ca      *x509.Certificate
	private *rsa.PrivateKey
	caName  = "mitm.crt"
	priName = "pri.pem"
)

func LoadCert() error {
	c, err := loadRootCA()
	if err != nil {
		return err
	}

	ca = c

	key, err := loadRootKey()
	if err != nil {
		return err
	}

	private = key

	return nil
}

// 加载根证书
func loadRootCA() (*x509.Certificate, error) {
	mitmData, err := os.ReadFile(caName)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(mitmData)

	return x509.ParseCertificate(block.Bytes)
}

// 加载根证书私钥
func loadRootKey() (*rsa.PrivateKey, error) {
	priData, err := os.ReadFile(priName)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(priData)

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func GenMITM() error {
	// 生成证书, 私钥
	caData, priData, err := newCert(domain, organization, validDuration)
	if err != nil {
		return err
	}

	// 写入新的私钥
	keyF, err := os.Create(priName)
	if err != nil {
		return err
	}
	defer keyF.Close()
	if _, err := keyF.Write(priData); err != nil {
		return err
	}

	// 写入新的mitm
	mitmF, err := os.Create(caName)
	if err != nil {
		return err
	}
	defer mitmF.Close()
	if _, err := mitmF.Write(caData); err != nil {
		return err
	}

	return nil
}

func newCert(domain string, organization string, validDuration time.Duration) ([]byte, []byte, error) {
	pub, prv, err := mitm.NewAuthority(domain, organization, validDuration)
	if err != nil {
		return nil, nil, err
	}
	prvData := bytes.NewBuffer([]byte{})
	if err = pem.Encode(prvData, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(prv)}); err != nil {
		return nil, nil, err
	}
	certData := bytes.NewBuffer([]byte{})
	if err = pem.Encode(certData, &pem.Block{Type: "CERTIFICATE", Bytes: pub.Raw}); err != nil {
		return nil, nil, err
	}

	return certData.Bytes(), prvData.Bytes(), nil
}
