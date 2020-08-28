package main

import (
	"fmt"
	cli "github.com/TharpHuang/gofar/gofar-cli/command"
	gcmd "github.com/TharpHuang/gofar/tools/os"
	"github.com/TharpHuang/gofar/tools/text"
)

const (
	VERSION = "0.0.1"
)

var (
	helpContent = text.TrimLeft(`
USAGE
    gofar COMMAND [ARGUMENT] [OPTION]
COMMAND
    gen        automatically generate go files of options type...
    help       show more information about a specified command
    build      cross-building go project for lots of platforms...
    version    show current binary version info...
	migrate    migrate the migration files of database... 
OPTION
    file name or some other things
ADDITIONAL
    Use 'gofar help COMMAND' or 'gofar COMMAND -h' for detail about a command
`)
)

func main() {

	command := gcmd.GetArg(1)

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
	default:
		fmt.Println("do you forget something?")
	}
}

// help shows more information for specified command.
func help(command string) {
	switch command {
	case "version":
		fmt.Print("help")
	default:
		fmt.Print(helpContent)
	}
}

func version() {
	fmt.Printf("Gofar CLI Tool: %s \n", VERSION)
}

func genProto() {

}
