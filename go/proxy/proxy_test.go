package proxy

import (
	"fmt"
	"os"
	"testing"
)

func TestServe(t *testing.T) {
	err := Serve()
	if err != nil {
		t.Fatal(err)
	}
}

func TestMartian(t *testing.T) {
	t.Run("localhost", func(t *testing.T) {
		err := Martian()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("proxy", func(t *testing.T) {

		b, err := ExportCrt()
		if err != nil {
			t.Fatal(err)
		}

		f, err := os.Create("mitm.crt")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = f.Write(b)
		_ = f.Close()
		fmt.Println("双击 mitm.crt , 存放到受信任的根证书颁发机构")

		ProxyServer = "http://127.0.0.1:10809"
		err = Martian()
		if err != nil {
			t.Fatal(err)
		}
	})
}
