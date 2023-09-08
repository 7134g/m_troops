package main

import (
	"path"
	"strings"
	"testing"
)

func TestName(t *testing.T) {

	dirName := strings.ReplaceAll(fileName, path.Ext(fileName), "")
	t.Log(dirName)

}
