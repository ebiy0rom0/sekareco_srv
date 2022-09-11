package infra

import (
	"fmt"
	"os"
	"sekareco_srv/domain/infra"
	"sekareco_srv/util"
)

// log level
type LogLevel int8

const (
	ERROR LogLevel = iota + 1
	WARN
	INFO
)

// logger instance
var Logger infra.Logger

type LogManager struct {
	e *os.File
	w *os.File
	i *os.File
}

func InitLogger() error {
	efp, err := os.OpenFile(errorLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	wfp, err := os.OpenFile(warnLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	ifp, err := os.OpenFile(infoLogFilePath(), os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	Logger = &LogManager{
		e: efp,
		w: wfp,
		i: ifp,
	}
	return nil
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
	return fmt.Sprintf("[%s]%s\n", Timer.NowDatetime(), err.Error())
}

func errorLogFilePath() string {
	return util.RootDir() + os.Getenv("LOG_PATH") + os.Getenv("ERROR_LOG_FILE_NAME")
}

func warnLogFilePath() string {
	return util.RootDir() + os.Getenv("LOG_PATH") + os.Getenv("WARN_LOG_FILE_NAME")
}

func infoLogFilePath() string {
	return util.RootDir() + os.Getenv("LOG_PATH") + os.Getenv("INFO_LOG_FILE_NAME")
}

var _ infra.Logger = &LogManager{}
