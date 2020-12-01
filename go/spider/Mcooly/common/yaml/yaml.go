package yaml

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"m_troops/go/spider/Mcooly/common/logs"
	"m_troops/go/spider/Mcooly/setting"
	"strconv"
	"time"
)

type MysqlParams struct {
	Mysql struct {
		User     string `yaml:"user"`
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
		DBName   string `yaml:"dbname"`
	}
	Timer struct {
		Hour   int `yaml:"hour"`
		Minute int `yaml:"minute"`
	}
	LogLevel   string `yaml:"loglevel"`
	SpiderTime struct {
		Defalut    string `yaml:"defalut"`
		Ccgp       string `yaml:"ccgp"`
		Hnbidding  string `yaml:"hnbidding"`
		Hnztb      string `yaml:"hnztb"`
		Regulatory string `yaml:"regulatory"`
		Serveplat  string `yaml:"serveplat"`
	}
	ProxyHost string `yaml:"proxyhost"`
}

func LoadYaml() {
	conf := new(MysqlParams)
	yamlFile, err := ioutil.ReadFile("env.yaml")

	if err != nil {
		logs.Exit("yamlFile.Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		logs.Exit("Unmarshal: %v", err)
	}
	setting.MYSQLCONNYAML = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		conf.Mysql.User,
		conf.Mysql.Password,
		conf.Mysql.Host,
		conf.Mysql.Port,
		conf.Mysql.DBName,
	)
	setting.TIMERHOUR = conf.Timer.Hour
	setting.TIMERMINUTE = conf.Timer.Minute

	switch conf.LogLevel {
	case "debug":
		setting.LOGLEVEL = logrus.DebugLevel
	case "info":
		setting.LOGLEVEL = logrus.InfoLevel
	case "error":
		setting.LOGLEVEL = logrus.ErrorLevel
	default:
		setting.LOGLEVEL = logrus.DebugLevel
	}

	defalut, err := strconv.ParseInt(conf.SpiderTime.Defalut, 10, 64)
	if err != nil {
		logs.Exit("strconv.ParseInt ERROR: ", err, " VALUE: ", defalut)
	}
	setting.WAIT_TIME_DEFULT = time.Duration(defalut) * time.Second

	serveplat, err := strconv.ParseInt(conf.SpiderTime.Serveplat, 10, 64)
	if err != nil {
		logs.Exit("strconv.ParseInt ERROR: ", err, " VALUE: ", serveplat)
	}
	setting.WAIT_TIME_SERVEPLAT = time.Duration(serveplat) * time.Second

	if conf.ProxyHost != "" {
		setting.PROXY_HOST = conf.ProxyHost
	}
}
