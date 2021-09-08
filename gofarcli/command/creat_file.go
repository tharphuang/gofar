package command

import (
	"fmt"
	"log"
	"os"

	"github.com/TharpHuang/gofar/tools/text"
	"github.com/spf13/afero"
)

var (
	warningInit = text.TrimLeft(`
Gofar build: no arguments input.
Use "gofar init [arguments]"
The arguments are:
	[model]	the name of model what you want to init.  
`)
	builder afero.Afero
)

func InitModel(params string) {
	if params == "" {
		fmt.Println(warningInit)
		return
	}
	modelName := params
	fmt.Println(modelName)
}

func CreateFileSystem(projectName string) {
	rootPath, err := os.Getwd()
	projectPath := rootPath + projectName
	_, err = afero.DirExists(builder, projectPath)
	if err != nil {
		fmt.Println(err)
	}
	err = builder.Mkdir(projectPath, os.ModeDir)
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Project init success!")
}
