package proxy

import (
	"github.com/google/martian"
	"net"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init("cert.sqlite3")
	os.Exit(m.Run())
}

func TestCert(t *testing.T) {
	mc, err := GlobalCert.GetMITMConfig()
	if err != nil {
		t.Fatal(err)
	}

	proxy := martian.NewProxy()
	proxy.SetRequestModifier(&Skip{})
	proxy.SetMITM(mc)

	listener, err := net.Listen("tcp", ":1080")
	if err != nil {
		t.Fatal(err)
	}

	err = proxy.Serve(listener)
	if err != nil {
		t.Fatal(err)
	}
}
