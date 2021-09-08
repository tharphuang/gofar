package os

import (
	"os"

	gregex "github.com/TharpHuang/gofar/tools/text"
)

var (
	defaultParsedArgs     = make([]string, 0)
	defaultParsedOptions  = make(map[string]string)
	defaultCommandFuncMap = make(map[string]func())
)

func doInit() {
	if len(defaultParsedArgs) > 0 {
		return
	}

	// Parsing os.Args with default algorithm.
	// The option should use '=' to separate its name and value in default.
	for _, arg := range os.Args {
		array, _ := gregex.MatchString(`^\-{1,2}([\w\?\.\-]+)={0,1}(.*)$`, arg)
		if len(array) == 3 {
			defaultParsedOptions[array[1]] = array[2]
		} else {
			defaultParsedArgs = append(defaultParsedArgs, arg)
		}
	}
}

func GetArg(index int, def ...string) string {
	doInit()

	if index < len(defaultParsedArgs) {
		return defaultParsedArgs[index]
	}

	if len(def) > 0 {
		return def[0]
	}

	return ""
}

func ContainsOpt(name string, def ...string) bool {
	doInit()
	_, ok := defaultParsedOptions[name]
	return ok
}
