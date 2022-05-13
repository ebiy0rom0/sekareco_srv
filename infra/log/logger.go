package logger

import (
	"fmt"
	"os"
)

const (
	LOG_LEVEL_ERROR = iota
	LOG_LEVEL_WARN
	LOG_LEVEL_INFO
)

var errorLogFile *os.File
var warnLogFile *os.File
var infoLogFile *os.File

func InitLogger() {
	fp, err := os.OpenFile(errorLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	errorLogFile = fp

	fp, err = os.OpenFile(warnLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	warnLogFile = fp

	fp, err = os.OpenFile(infoLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	infoLogFile = fp
}

func CleanupLogger() error {
	errorLogFile.Close()
	warnLogFile.Close()
	infoLogFile.Close()

	return dropLogFile()
}

// for debug
func dropLogFile() error {
	return os.RemoveAll(os.Getenv("LOG_PATH"))
}

func Logging(level int, msg string) {
	switch level {
	case LOG_LEVEL_ERROR:
		loggingError(msg)
	case LOG_LEVEL_WARN:
		loggingWarn(msg)
	case LOG_LEVEL_INFO:
		loggingInfo(msg)
	}
}

func loggingError(msg string) {
	fmt.Fprintln(errorLogFile, msg)
}

func loggingWarn(msg string) {
	fmt.Fprintln(warnLogFile, msg)
}

func loggingInfo(msg string) {
	fmt.Fprintln(infoLogFile, msg)
}

func errorLogFilePath() string {
	return os.Getenv("LOG_PATH") + os.Getenv("ERROR_LOG_FILE_NAME")
}

func warnLogFilePath() string {
	return os.Getenv("LOG_PATH") + os.Getenv("WARN_LOG_FILE_NAME")
}

func infoLogFilePath() string {
	return os.Getenv("LOG_PATH") + os.Getenv("INFO_LOG_FILE_NAME")
}
