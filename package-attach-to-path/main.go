package main

import (
	"flag"
	"os"
	"path/filepath"
)

var (
	generateInit        = flag.Bool("generate_init", false, "generate init file")
	generateRegisterMap = flag.Bool("generate_register_map", false, "generate register file")
	_currentPackageName = flag.String("package", "", "generate register file from name")
	_currentPackagePath = flag.String("path", "", "generate register file from path")
)

func init() {
	if !flag.Parsed() {
		flag.Parse()
	}
}

func main() {
	if *generateRegisterMap {
		generateRegisterMapFile()
	}

	if *generateInit {
		generateInitFiles()
	}
}

func generateInitFiles() {
	path := filepath.Join(initTemplateVars().CurrentPath, "package-probe")
	err := os.MkdirAll(path, 0755)
	if err != nil {
		panic(err)
	}
	toFile(`package probe

import (
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

type __PROBE___ struct {}
var __PROBE__ __PROBE___
var RootPackage string
var RootPath string

func init() {
	var InstancePackage, InstancePath string
	InstancePackage = reflect.TypeOf(__PROBE__).PkgPath()
	RootPackage = strings.ReplaceAll(filepath.Dir(InstancePackage), "\\", "/")
	_, InstancePath, _, _ = runtime.Caller(0)
	InstancePath = filepath.Dir(InstancePath)
	RootPath = filepath.Dir(InstancePath)
}
`, filepath.Join(initTemplateVars().CurrentPath, "package-probe/init.go"), initTemplateVars())

	toFile(`package probe

import (
	"fmt"
	"testing"
)

func TestPackage(t *testing.T) {
	fmt.Println(RootPackage, RootPath)
}
`, filepath.Join(initTemplateVars().CurrentPath, "package-probe/probe_test.go"), initTemplateVars())

	toFile(`// Code generated attach-file.go DO NOT EDIT
package {{.PackageName}}

import (
	_ "{{.Package}}/package-probe"
)

`, filepath.Join(initTemplateVars().CurrentPath, "init-package-probe.go"), initTemplateVars())
}

func generateRegisterMapFile() {
	toFile(`// Code generated minimum-attach-file.go DO NOT EDIT
package {{.PackageName}}

import instance "github.com/Myriad-Dreamin/go-magic-package/instance"

func init() {
	instance.Register("{{.Package}}", `+"`"+`{{.CurrentPath}}`+"`"+`)
}
`, initTemplateVars().CurrentPath+"/register-instance-map.go", initTemplateVars())
}
