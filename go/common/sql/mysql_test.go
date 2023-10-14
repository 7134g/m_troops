package model

import "testing"

func TestDatabase(t *testing.T) {
	Database("root:123456@tcp(127.0.0.1:3306)/hnspider?charset=utf8")
	_ = DB.Close()
}
