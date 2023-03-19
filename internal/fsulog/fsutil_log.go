package fsulog

import (
	"os"

	"github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var (
	debugLogger *logrus.Logger
)

func init() {
	initLogging()
}

func initLogging() {
	debugLogger = &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}
}

func GetLogger() *logrus.Logger {
	return debugLogger
}
