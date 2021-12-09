package command

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"

	"github.com/AlecAivazis/survey/v2"
)

/**
 生成项目目录：
	- project
		- someservice
			- cmd
			- middles
			- svc
			- api
		- otherservice
		- main.go
*/

var InitCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i"},
	Short:   "Create a new framework base our model of your project",
	Example: `  gofar initialization a new project`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(path)
		if len(args) == 0 {
			ExecError(cmd, args, errors.New("unknown project name"))
		}

		err := initialization(args[0])
		if err != nil {
			ExecError(cmd, args, errors.New("unknown project name"))
		}
	},
}

var tpl = `package {{.Package}}

import "fmt"

func main(){
	// todo: add some service or middles
    fmt.Println("hello")
}
`

type Project struct {
	Name string
	Path string
}

func (p *Project) NewProject(dir, name string) error {
	fmt.Println(name)
	p.Name = name
	p.Path = path.Join(dir, name)

	if _, err := os.Stat(p.Path); !os.IsNotExist(err) {
		fmt.Printf("%s already exists \n", p.Path)
		override := false
		option := &survey.Confirm{
			Message: "Do you want remove it and recreate",
			Help:    "Delete the existed project and create a new",
		}
		surveyErr := survey.AskOne(option, &override)
		if surveyErr != nil {
			return surveyErr
		}
		if !override {
			return err
		}
		os.RemoveAll(p.Path)
	}

	err := os.MkdirAll(p.Path, 0755)
	if err != nil {
		return err
	}

	var dest = path.Join(p.Path, strings.ToLower("main")+".go")
	generateGoFile(tpl, "main", dest)
	return nil
}

func generateGoFile(temp, pack, dest string) {
	tt := template.Must(template.New("queue").Parse(temp))

	var (
		file *os.File
		err  error
	)

	if file, err = os.Create(dest); err != nil {
		if !os.IsExist(err) {
			fmt.Printf("Could not create %s: %s (skip)\n", dest, err)
		}
		_ = os.Remove(dest)
	}
	tmpMain := map[string]string{
		"Package": pack,
	}
	_ = tt.Execute(file, tmpMain)
	_ = file.Close()
}

func initialization(name string) error {
	p := Project{}
	pwd, _ := os.Getwd()
	err := p.NewProject(pwd, name)
	return err
}
