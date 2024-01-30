package proxy

import (
	"testing"
)

func TestMartian(t *testing.T) {
	SetServeProxyAddress("http://127.0.0.1:7890", "", "")
	OpenCert()
	if err := Martian(); err != nil {
		t.Fatal(err)
	}

}
