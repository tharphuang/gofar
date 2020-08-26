package main

import ""

const (
	VERSION = "0.0.1"
)

var (
	helpContent = string(`
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