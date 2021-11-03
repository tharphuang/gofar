package xlog

var DefaultLogger = NewLogger()

func Info(args ...interface{}){
	DefaultLogger.Info(args)
}
