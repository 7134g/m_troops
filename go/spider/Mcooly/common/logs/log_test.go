package logs

import (
	"errors"
	"testing"
)

func TestInitTheWorldLog(t *testing.T) {
	InitTheWorldLog()
	Log.Info("aaaaaa")
	Log.Error("bbbbbb")
	Log.Debug("cccccccccc")
	Log.Warning("sssssssssssss")
}

func TestError(t *testing.T) {
	InitTheWorldLog()
	Errors("aaa", errors.New("xxxxx"))
	Errors(errors.New("xxxxx"), "aaa")
	Errors(errors.New("xxxxx"), "aaa", 1)
	Errors(errors.New("xxxxx"), "aaa", nil, 1)
}
