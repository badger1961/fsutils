package fsulog

import (
	"fmt"
	"log"
	"os"
)

var (
	debugLogger   *log.Logger
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func init() {
	err := initLogging()
	if err != nil {
		fmt.Printf("Fatal event, logging initiating failed")
		os.Exit(1)
	}

}

func initLogging() error {

	debugLogger = log.New(os.Stdout, "DEBUG:", log.Ldate|log.Ltime|log.Lshortfile|log.Lmicroseconds)
	warningLogger = log.New(os.Stdout, "WARN:", log.Ldate|log.Ltime|log.Lshortfile|log.Lmicroseconds)
	infoLogger = log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime|log.Lshortfile|log.Lmicroseconds)
	errorLogger = log.New(os.Stdout, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile|log.Lmicroseconds)
	infoLogger.Println("FSUtils Logging has been started")
	return nil
}

func GetDebugLogger() *log.Logger {
	return debugLogger
}

func GetWarnLogger() *log.Logger {
	return warningLogger
}

func GetInfoLogger() *log.Logger {
	return infoLogger
}

func GetErrorLogger() *log.Logger {
	return errorLogger
}
