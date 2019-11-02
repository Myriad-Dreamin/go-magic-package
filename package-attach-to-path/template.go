package main

import (
	"github.com/Myriad-Dreamin/go-parse-package"
	"github.com/Myriad-Dreamin/market/lib/sugar"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func initTemplate(t string) *template.Template {
	var resultTemplate = template.New("resultTemplate")

	var err error
	resultTemplate, err = resultTemplate.Parse(t)
	if err != nil {
		panic(err)
	}
	return resultTemplate
}

type TemplateVars struct {
	PackageName, Package, CurrentPath string
}

var vars *TemplateVars

func initTemplateVars() *TemplateVars {
	return vars
}

func toFile(t, path string, vars *TemplateVars) {
	var resultTemplate = initTemplate(t)
	sugar.WithWriteFile(func(outputFile *os.File) {
		err := resultTemplate.Execute(outputFile, vars)
		if err != nil {
			panic(err)
		}
	}, path)
}

func init() {
	var currentPath, currentPackageName, currentPackagePath string
	var err error
	currentPath, err = filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	if len(*_currentPackagePath) != 0 {
		currentPackagePath = *_currentPackagePath
	} else {
		goPath := os.Getenv("GOPATH")
		if len(goPath) != 0 {
			currentPackagePath, err = filepath.Rel(filepath.Join(goPath, "src"), currentPath)
			if err != nil {
				panic(err)
			}
		} else {
			//find go.mod
			panic("package path undefined")
		}
	}
	currentPackagePath = strings.ReplaceAll(currentPackagePath, "\\", "/")
	if len(*_currentPackageName) != 0 {
		currentPackageName = *_currentPackageName
	} else {
		currentPackageName, err = parser.ParsePackageName(currentPath)
		if err != nil {
			currentPackageName = "main"
		}
	}

	vars = &TemplateVars{
		PackageName: currentPackageName,
		Package:     currentPackagePath,
		CurrentPath: currentPath,
	}
}
