package proxy

import (
	"bytes"
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"github.com/google/martian/mitm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"path/filepath"
	"time"
)

var (
	_db *gorm.DB
)

func Init(dbname string) {
	var (
		err error
	)
	dbname, err = filepath.Abs(dbname)
	_db, err = gorm.Open(sqlite.Open(dbname), &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.Default.LogMode(logger.Error),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatalln("连接数据库失败", err)
	}
	sqlDB, err := _db.DB()
	if err != nil {
		log.Fatalln("连接数据库失败", err)
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	c := &Cert{}
	_ = _db.AutoMigrate(c)
	c.init()
}

var (
	Domain              = "proxy"     //配置API域名
	Organization        = "Authority" //MITM证书的组织名称
	CAValidDay   uint32 = 3650        //MITM证书有效天数（单位天）
	Validity     uint32 = 1           //MITM证书有效的时间窗口（单位小时）

	GlobalCert *Cert // 全局变量证书对象
)

type Cert struct {
	ID         uint   `gorm:"column:id" json:"-"`
	PrivateKey string `gorm:"column:private_key" json:"private_key"`
	PublicKey  string `gorm:"column:public_key" json:"public_key"`
}

func (c *Cert) init() {
	var _cert *Cert
	var err error
	if _cert, err = c.Get(); errors.Is(err, gorm.ErrRecordNotFound) {
		_cert, err = newCert(Domain, Organization, time.Duration(CAValidDay)*time.Hour*24)
		if err != nil {
			return
		}
		if err := _cert.Create(); err != nil {
			return
		}
	}

	GlobalCert = _cert
}

func (c *Cert) Create() error {
	return _db.Create(c).Error
}

func (c *Cert) Get() (*Cert, error) {
	err := _db.First(c).Error
	return c, err
}

// GetMITMConfig 获取证书配置对象
func (c *Cert) GetMITMConfig() (*mitm.Config, error) {
	ca, pri, err := genMITMConfig(c)
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
func (c *Cert) ExportCrt() ([]byte, error) {
	ca, _, err := genMITMConfig(c)
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
