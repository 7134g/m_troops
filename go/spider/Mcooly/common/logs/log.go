package logs

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	Log  *logrus.Logger
	once sync.Once
)

type MyFormatter struct{}

func (s *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006/01/02 15:04:05")
	msg := fmt.Sprintf("%s [%s] %s\n", timestamp, strings.ToUpper(entry.Level.String()), entry.Message)
	return []byte(msg), nil
}

func InitTheWorldLog() {
	once.Do(func() {
		Log = logrus.New()
		Log.SetLevel(setting.LOGLEVEL)
		Log.SetFormatter(new(MyFormatter))
		Log.SetOutput(os.Stdout)
	})
}

func Errors(errs ...interface{}) {
	if errs[0] == nil {
		return
	}
	var nError []interface{}
	for i := 0; i < len(errs); i++ {
		if errs[i] != nil {
			nError = append(nError, errs[i])
			nError = append(nError, " ")
		}
	}
	Log.Error(nError...)
}

func Exit(errs ...interface{}) {
	Log.Fatalln(errs...)
}
