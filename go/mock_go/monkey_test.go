package mock_go

import (
	"bou.ke/monkey"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestMyFunc(t *testing.T) {
	// 对 GetInfoByUID 进行打桩
	// 无论传入的uid是多少，都返回 &UserInfo{Name: "liwenzhou"}, nil
	monkey.Patch(GetInfoByUID, func(int64) (*UserInfo, error) {
		return &UserInfo{Name: "monkey"}, nil
	})

	ret := MyFunc(123)
	if !strings.Contains(ret, "monkey") {
		t.Fatal(ret)
	}
	t.Log(ret)
	monkey.UnpatchAll()
	t.Log(MyFunc(123))
}

func TestNet1(t *testing.T) {
	var d *net.Dialer // Has to be a pointer to because `Dial` has a pointer receiver
	monkey.PatchInstanceMethod(reflect.TypeOf(d), "Dial", func(_ *net.Dialer, _, _ string) (net.Conn, error) {
		return nil, fmt.Errorf("no dialing allowed")
	})
	_, err := http.Get("http://google.com")
	fmt.Println(err) // Get http://google.com: no dialing allowed
}

func TestNet(t *testing.T) {
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(http.DefaultClient), "Get", func(c *http.Client, url string) (*http.Response, error) {
		guard.Unpatch()
		defer guard.Restore()

		if !strings.HasPrefix(url, "https://") {
			return nil, fmt.Errorf("only https requests allowed")
		}

		return c.Get(url)
	})

	_, err := http.Get("http://google.com")
	fmt.Println(err) // only https requests allowed
	//guard.Unpatch()
	//resp, err := http.Get("https://google.com")
	//fmt.Println(resp.Status, err) // 200 OK <nil>
}

func TestMyClient_Post(t *testing.T) {
	var m *MyClient
	var guard *monkey.PatchGuard
	guard = monkey.PatchInstanceMethod(reflect.TypeOf(m), "Post", func(c *MyClient, url string, body []byte, header map[string]interface{}) (*http.Response, error) {
		guard.Unpatch()
		defer guard.Restore()

		if strings.Contains(url, "x") {
			return nil, errors.New("error ok")
		}

		return c.Post(url, nil, nil)
	})

	x, err := m.Post("x", nil, nil)
	fmt.Println(x, err)
}

func TestFmtPrintln_Monkey(t *testing.T) {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		b := []interface{}{
			"new",
		}
		b = append(b, a...)
		return fmt.Fprintln(os.Stdout, b...)
	})

	fmt.Println("x")
}
