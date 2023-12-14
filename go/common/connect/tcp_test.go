package connect

import (
	"log"
	"testing"
	"time"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestNewDial(t *testing.T) {
	d, err := NewDial("tcp", "127.0.0.1:2222", time.Second*10)
	if err != nil {
		log.Fatal(err)
	}

	for {
		data := make([]byte, 1024)
		n, err := d.Read(data)
		if err != nil {
			log.Fatal(string(data[:n]), err)
		}

		log.Println(string(data[:n]))
	}
}
