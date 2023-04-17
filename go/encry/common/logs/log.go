package logs

import (
	"encry/config"
	"io/ioutil"
	"log"
)

func Load() {
	if config.LOGSILENT == 1 {
		log.SetOutput(ioutil.Discard)
	}
}

func Info(v ...interface{}) {
	newErr := make([]interface{}, 0)
	newErr = append(newErr, "[Info] ")
	newErr = append(newErr, v...)
	log.Println(newErr...)
}

func Error(v ...interface{}) {
	newErr := make([]interface{}, 0)
	newErr = append(newErr, "[ERROR] ")
	newErr = append(newErr, v...)
	log.Println(newErr...)
}

func Exit(v ...interface{}) {
	log.Fatalln(v...)
}
