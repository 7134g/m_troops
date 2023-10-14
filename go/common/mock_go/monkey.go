package mock_go

import (
	"fmt"
	"net/http"
)

type UserInfo struct {
	UID  int64
	Name string
}

func MyFunc(uid int64) string {
	u, err := GetInfoByUID(uid)
	if err != nil {
		return "welcome"
	}

	// 这里是一些逻辑代码...

	return fmt.Sprintf("hello %s\n", u.Name)
}

func GetInfoByUID(uid int64) (*UserInfo, error) {
	return &UserInfo{Name: "lxc", UID: uid}, nil
}

type MyClient struct {
}

func (m *MyClient) Post(url string, _ []byte, _ map[string]interface{}) (*http.Response, error) {
	return http.Post(url, "application/json", nil)
}
