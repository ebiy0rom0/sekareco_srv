package logger

import (
	"os"
	infraDomain "sekareco_srv/domain/infra"
	"sekareco_srv/util"

	"github.com/rs/zerolog"
)

type fileLogger struct {
	i zerolog.Logger
	w zerolog.Logger
	e zerolog.Logger
}

// NewFileLogger returns new fileLogger.
// fileLogger is implements of infra.ILogger.
func NewFileLogger() (*fileLogger, error) {
	ifp, err := os.OpenFile(logLocate() + "info.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	wfp, err := os.OpenFile(logLocate() + "warn.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	efp, err := os.OpenFile(logLocate() + "error.log", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &fileLogger{
		i: zerolog.New(ifp),
		w: zerolog.New(wfp),
		e: zerolog.New(efp),
	}, nil
}

// Error writes info message into info log file.
func (l *fileLogger) Info(err error) {
	l.i.Info().Msgf("%+v", err)
}

// Error writes warning message into warn log file.
func (l *fileLogger) Warn(err error) {
	l.w.Warn().Msgf("%+v", err)
}

// Error writes error message into error log file.
func (l *fileLogger) Error(err error) {
	l.e.Error().Msgf("%+v", err)
}

// logLocate returns a string of the log file path.
func logLocate() string {
	return util.RootDir() + os.Getenv("LOG_PATH")
}

var _ infraDomain.ILogger = (*fileLogger)(nil)