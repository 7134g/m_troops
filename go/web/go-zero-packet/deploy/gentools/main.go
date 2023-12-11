package main

import (
	"flag"
)

var (
	mysqlLink = "root:mysql@tcp(127.0.0.1:3306)/blog"

	packName   = "gen"
	mysqlDbDir = "./manage/mysqldb/query"
	mongoDbDir = "./manage/mongodb/model"

	mysqlTableNames = ""
	mongoTableNames = ""
	templateDir     = ""
)

func main() {
	flag.StringVar(&packName, "p", packName, "pack name, default: gen")
	flag.StringVar(&mysqlTableNames, "mysql", mysqlTableNames, "mysql table names, 逗号分割")
	flag.StringVar(&mongoTableNames, "mongo", mongoTableNames, "mongo table names, 逗号分割")
	flag.StringVar(&templateDir, "t", templateDir, "模板文件所在地")
	flag.Parse()

	// -t ./template -mysql work_object_attr
	// -t ./template -mongo UserActivity
	run()
}

func run() {

	mysql := MysqlInfo{
		TableNames:         mysqlTableNames,
		Link:               mysqlLink,
		PackName:           packName,
		Dir:                mysqlDbDir,
		TemplateDir:        templateDir,
		TemplatePath:       mysqlTemplate,
		TableField:         map[string][]mysqlField{},
		QueryStructNameMap: map[string]string{},
		ModelStructNameMap: map[string]string{},
	}
	mysql.GenerateModel()

	mongo := MongoConfig{
		TableNames:  mongoTableNames,
		TemplateDir: templateDir,
		Dir:         mongoDbDir,
		PackName:    packName,
		Table:       mongoTemplate,
	}
	mongo.GenerateModel()
}
