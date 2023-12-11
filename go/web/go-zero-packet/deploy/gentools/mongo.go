package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

type MongoConfig struct {
	TableNames string

	PackName    string
	TemplateDir string
	Dir         string

	Table map[string][]string
}

func (m *MongoConfig) Check() {
	if err := exec.Command("goctl", "--help").Run(); err != nil {
		log.Fatal(err)
	}
}

func (m *MongoConfig) GenerateModel() {
	if m.TableNames == "" {
		log.Println("mongoTableNames is empty")
		return
	}

	tableNames := strings.Split(m.TableNames, ",")

	for _, name := range tableNames {
		args := []string{
			"model",
			"mongo",
			"--type",
			name,
			"--dir",
			m.Dir,
			"--style",
			"go_zero",
		}

		if m.TemplateDir != "" {
			args = append(args, "--home")
			args = append(args, m.TemplateDir)
		}

		fmt.Println(append([]string{"goctl"}, args...))
		if err := exec.Command("goctl", args...).Run(); err != nil {
			fmt.Println(name, err)
		}
	}

	m.GenerateManage()
}

func (m *MongoConfig) GenerateManage() {
	if m.TableNames == "" {
		log.Println("mongoTableNames is empty")
		return
	}

	modelPackPath := path.Join(m.PackName, m.Dir, "../", "model")
	headTemplate := TemplateObject{
		ModelPackPath: modelPackPath,
	}

	tableNames := strings.Split(m.TableNames, ",")
	bodyTemplate := make([]TemplateObject, 0)
	for _, name := range tableNames {
		bodyTemplate = append(bodyTemplate, TemplateObject{
			Type: name,
		})
	}

	for dst, templateList := range m.Table {
		tg := templateGroup{
			targets:   make([]target, 0),
			body:      bytes.NewBuffer(nil),
			writePath: dst,
		}

		if _, err := os.Stat(dst); err != nil {
			tg.rewrite = true
			if err := os.MkdirAll(filepath.Dir(dst), os.ModeDir); err != nil {
				log.Fatal(err)
			}
		}
		for i := 0; i < len(templateList); i++ {
			templateList[i] = filepath.Join(m.TemplateDir, templateList[i])
			if _, err := os.Stat(templateList[i]); err != nil {
				log.Println(err)
				continue
			}
		}

		switch {
		case strings.Contains(dst, "query.go"):
			tg.rewrite = true
			tg.targets = append(tg.targets,
				target{
					srcTemplatePath: templateList[0],
					data:            headTemplate,
				},
				target{
					srcTemplatePath: templateList[1],
					data:            bodyTemplate,
				},
			)
		case strings.Contains(dst, "config.go"):
			tg.targets = append(tg.targets, target{
				srcTemplatePath: templateList[0],
				data:            bodyTemplate,
			})
		}

		if err := tg.parse(); err != nil {
			log.Fatal(err)
		}
		FmtFile(dst)
	}

}
