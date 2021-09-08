package discription

import "github.com/TharpHuang/gofar/tools/text"

var (
	HelpContent = text.TrimLeft(`
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
