package proxy

import "testing"

func TestServe(t *testing.T) {
	err := Serve()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMartian(t *testing.T) {
	err := Martian()
	if err != nil {
		t.Fatal(err)
	}
}
