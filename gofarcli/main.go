package main

import (
	"github.com/TharpHuang/gofar/gofarcli/command"
)

func main() {
	command.Execute()
	/*command := gcmd.GetArg(1)

	if gcmd.ContainsOpt("h") && command != "" {
		help(command)
		return
	}

	switch command {
	case "help":
		help(gcmd.GetArg(2))
	case "version":
		version()
	case "gen":
		cli.GenerateFile(gcmd.GetArg(2), gcmd.GetArg(3))
	case "build":
		cli.BuildFile(gcmd.GetArg(2))
	case "migrate":
		cli.Migrate(gcmd.GetArg(2))
	case "init":
		cli.CreateFileSystem(gcmd.GetArg(2))
	default:
		fmt.Println("do you forget something?")
	}*/
}
