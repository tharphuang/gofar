package main

import (
	"fmt"
	cmd "gofar/gofar-cli/command"
	gcmd "gofar/tools/os"
	gstr "gofar/tools/text"
)

const (
	VERSION = "0.0.1"
)

var (
	helpContent = gstr.TrimLeft(`
USAGE
    gf COMMAND [ARGUMENT] [OPTION]
COMMAND
    env        show current Golang environment variables
    get        install or update GF to system in default...
    gen        automatically generate go files for ORM models...
    mod        extra features for go modules...
    run        running go codes with hot-compiled-like feature...
    init       initialize an empty GF project at current working directory...
    help       show more information about a specified command
    pack       packing any file/directory to a resource file, or a go file...
    build      cross-building go project for lots of platforms...
    docker     create a docker image for current GF project...
    swagger    swagger feature for current project...
    update     update current gf binary to latest one (might need root/admin permission)
    install    install gf binary to system (might need root/admin permission)
    version    show current binary version info
OPTION
    -y         all yes for all command without prompt ask 
    -?,-h      show this help or detail for specified command
    -v,-i      show version information
ADDITIONAL
    Use 'gf help COMMAND' or 'gf COMMAND -h' for detail about a command, which has '...' 
    in the tail of their comments.
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
		cmd.GenerateFile(gcmd.GetArg(2), gcmd.GetArg(3))
	case "build":
		cmd.BuildFile(gcmd.GetArg(2))
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
	fmt.Println(`GoFar CLI Tool %s`, VERSION)
}

func genProto() {

}
