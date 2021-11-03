package xlog

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

type ALog struct {
	file *os.File
	log  *logrus.Logger
}

func NewLogger() *ALog {
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Errorf(err.Error())
	}
	logFileName := "api.log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			fmt.Println(err.Error())
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	// 实例化
	var l = ALog{
		log:  logrus.New(),
		file: src,
	}

	l.log = logrus.New()

	// 设置输出
	l.log.Out = io.MultiWriter(src, os.Stdout)

	// 设置日志级别
	l.log.SetLevel(logrus.DebugLevel)

	// 设置日志格式
	l.log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	return &l
}

func (l *ALog) msg() {

}

func (l *ALog) Info(args ...interface{}) {
	l.log.Infof("%s", args)
}

func (l ALog) Error(args ...interface{}) {
	l.log.Errorf("%s", args)
}