package logger

import (
	"fmt"
	"os"
	"sekareco_srv/domain/infra"
	"sekareco_srv/infra/timer"
)

// log level
const (
	LOG_LEVEL_ERROR = iota
	LOG_LEVEL_WARN
	LOG_LEVEL_INFO
)

// logger instance
var Logger infra.Logger

type LogManager struct {
	e *os.File
	w *os.File
	i *os.File
}

func InitLogger() {
	m := new(LogManager)
	fp, err := os.OpenFile(errorLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.e = fp

	fp, err = os.OpenFile(warnLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.w = fp

	fp, err = os.OpenFile(infoLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	m.i = fp

	Logger = m
}

// for debug
func DropLogFile() {
	os.Remove(errorLogFilePath())
	os.Remove(warnLogFilePath())
	os.Remove(infoLogFilePath())
}

func (l *LogManager) Error(err error) {
	l.e.WriteString(logFormat(err))
}

func (l *LogManager) Warn(err error) {
	l.w.WriteString(logFormat(err))
}

func (l *LogManager) Info(err error) {
	l.i.WriteString(logFormat(err))
}

func logFormat(err error) string {
	return fmt.Sprintf("[%s]%s\n", timer.Timer.NowDatetime(), err.Error())
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
