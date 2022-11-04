package main

import (
	olog "log"
	"os"
	"strings"
	"time"
)

type sliceValue []string

func NewSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

func (s *sliceValue) Get() interface{} {
	return []string(*s)
}

func (s *sliceValue) String() string {
	return strings.Join([]string(*s), ",")
}

type Logger struct {
	status int
}

func (l *Logger) Fatalln(v ...interface{}) {
	if l.status == 1 {
		olog.Fatalln(v...)
	}
}

func (l *Logger) Println(v ...interface{}) {
	if l.status == 1 {
		olog.Println(v...)
	}
}

func (l *Logger) Printf(format string, v ...interface{}) {
	if l.status == 1 {
		olog.Printf(format, v...)
	}
}

func (l *Logger) SetLocalFile() {
	olog.SetOutput(openLog())
}

func openLog() *os.File {
	//args := os.Args
	//argc := len(os.Args)
	var logFileError error
	var logFile *os.File
	if LOGTPYE != "" {
		//timeStr := time.Now().Format("2006_01_02_15_04_05") // "2006-01-02 15:04:05"
		//logPath := LOGTPYE + "/" + timeStr + "-" + address1 + "_" + address2 + "-" + address3 + "_" + address4 + ".log"
		timeStr := time.Now().Format("2006_01_02") // "2006-01-02 15:04:05"
		logPath := LOGTPYE + "/" + timeStr + ".log"
		logPath = strings.Replace(logPath, `\`, "/", -1)
		logPath = strings.Replace(logPath, "//", "/", -1)
		logFile, logFileError = os.OpenFile(logPath, os.O_APPEND|os.O_CREATE, 0666)
		if logFileError != nil {
			log.Fatalln("[x] log file path error.", logFileError.Error())
		}
		log.Println("[√]", "open test log file success. path:", logPath)
	}
	return logFile
}
