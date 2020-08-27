package command

import (
	"fmt"
	"os"
	"os/exec"
)

func BuildFile(fileType, fileName string) {
	switch fileType {
	case "proto":
		buildProtoFile()
	default:
		fmt.Println("No such file type, please try again.")
	}

}

func buildProtoFile() {
	cmd, err := script("../real_deploy.sh")
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
