package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"text/template"
)

// 模版文件：生成go的路径文件
var (
	mongoTemplate = map[string][]string{
		"./manage/mongodb/query/query.go": []string{
			"additional/mongo_manage_head.tpl",
			"additional/mongo_manage_body.tpl",
		},
		"./manage/mongodb/model/config.go": []string{
			"additional/config.tpl",
		},
	}

	mysqlTemplate = map[string]string{
		"additional/mysql.tpl": "./manage/mysqldb/query/*.go",
	}
)

type TemplateObject struct {
	ModelPackPath string

	Type      string
	LowerType string

	Switch          string
	QueryStructName string
	ModelStructName string
}

func parseFile(templatePath, writePath string, data any) error {
	rf, err := os.Open(templatePath)
	if err != nil {
		return err
	}
	b, err := io.ReadAll(rf)
	if err != nil {
		return err
	}
	rf.Close()

	tmpl, err := template.New("gen").Parse(string(b))
	if err != nil {
		panic(err)
	}

	wf, err := os.Create(writePath)
	if err != nil {
		return err
	}
	defer wf.Close()
	err = tmpl.Execute(wf, data)
	if err != nil {
		return err
	}

	return nil
}

type templateGroup struct {
	targets []target

	body *bytes.Buffer

	rewrite   bool
	writePath string
}

type target struct {
	srcTemplatePath string
	data            any
}

func (c *templateGroup) parse() error {
	if !c.rewrite {
		return nil
	}

	for _, t := range c.targets {
		rf, err := os.Open(t.srcTemplatePath)
		if err != nil {
			return err
		}
		b, err := io.ReadAll(rf)
		if err != nil {
			return err
		}
		rf.Close()

		tmpl, err := template.New("gen").Parse(string(b))
		if err != nil {
			panic(err)
		}
		err = tmpl.Execute(c.body, t.data)
		if err != nil {
			return err
		}
	}

	wf, err := os.Create(c.writePath)
	if err != nil {
		return err
	}
	defer wf.Close()
	if _, err := wf.Write(c.body.Bytes()); err != nil {
		log.Fatal(err)
	}

	return nil
}
