package proxy

import (
	"bytes"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"github.com/google/martian/mitm"
	"time"
)

var (
	Domain              = "proxy"     //配置API域名
	Organization        = "Authority" //MITM证书的组织名称
	CAValidDay   uint32 = 3650        //MITM证书有效天数（单位天）
	//Validity     uint32 = 1           //MITM证书有效的时间窗口（单位小时）

	GlobalCert *Cert // 全局变量证书对象
)

type Cert struct {
	PrivateKey string `gorm:"column:private_key" json:"private_key"`
	PublicKey  string `gorm:"column:public_key" json:"public_key"`
}

func CertReload() {
	var err error
	GlobalCert, err = newCert(Domain, Organization, time.Duration(CAValidDay)*time.Hour*24)
	if err != nil {
		panic(err)
	}
}

// GetMITMConfig 获取证书配置对象
func GetMITMConfig() (*mitm.Config, error) {
	ca, pri, err := genMITMConfig(GlobalCert)
	if err != nil {
		return nil, err
	}
	mc, err := mitm.NewConfig(ca, pri)
	if err != nil {
		return nil, err
	}
	return mc, nil
}

// ExportCrt 导出CA证书
func ExportCrt() ([]byte, error) {
	ca, _, err := genMITMConfig(GlobalCert)
	if err != nil {
		return nil, err
	}
	content := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: ca.Raw,
	})

	return content, nil
}

func genMITMConfig(_cert *Cert) (*x509.Certificate, crypto.PrivateKey, error) {
	pub, err := base64.StdEncoding.DecodeString(_cert.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	pri, err := base64.StdEncoding.DecodeString(_cert.PrivateKey)
	if err != nil {
		return nil, nil, err
	}
	tlsc, err := tls.X509KeyPair(pub, pri)
	if err != nil {
		return nil, nil, err
	}
	x509c, err := x509.ParseCertificate(tlsc.Certificate[0])
	if err != nil {
		return nil, nil, err
	}
	return x509c, tlsc.PrivateKey, nil
}

func newCert(domain string, organization string, validDuration time.Duration) (*Cert, error) {
	pub, prv, err := mitm.NewAuthority(domain, organization, validDuration)
	if err != nil {
		return nil, err
	}
	prvKey := bytes.NewBuffer([]byte{})
	if err = pem.Encode(prvKey, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(prv)}); err != nil {
		return nil, err
	}
	pubKey := bytes.NewBuffer([]byte{})
	if err = pem.Encode(pubKey, &pem.Block{Type: "CERTIFICATE", Bytes: pub.Raw}); err != nil {
		return nil, err
	}

	return &Cert{PublicKey: base64.StdEncoding.EncodeToString(pubKey.Bytes()),
		PrivateKey: base64.StdEncoding.EncodeToString(prvKey.Bytes())}, nil
}
