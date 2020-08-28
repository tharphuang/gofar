package command

import (
	"fmt"
	"github.com/TharpHuang/gofar/tools/text"
	"os"
	"os/exec"
)

var warningBuild = text.TrimLeft(`
Gofar build: no arguments input.
Use "gofar build [arguments]"
The arguments are:
	proto	build protobuf files of go 
	migrate	build databases migrateion files of go`)

func BuildFile(jobType string) {
	if jobType == "" {
		fmt.Println(warningBuild)
		return
	}
	switch jobType {
	case "proto":
		buildProtoFile()
	default:
		fmt.Println(warningGen)
	}

}

func buildProtoFile() {
	cmd, err := script("./real_deploy.sh")
	if err != nil {
		panic(err)
	}
	fmt.Println(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}

	fmt.Println(string(output))

}

func script(scrip string) (*exec.Cmd, error) {
	shell := "/bin/sh"
	if other := os.Getenv("SHELL"); other != "" {
		shell = other
	}
	return exec.Command(shell, "-c", scrip+" generate protogen"), nil
}
