package logger

import (
	"os"
	"path/filepath"
	"sekareco_srv/domain/infra"
	"sekareco_srv/env"
	"sekareco_srv/util"

	"github.com/ebiy0rom0/errors"

	"github.com/rs/zerolog"
)

type fileLogger struct {
	i zerolog.Logger
	w zerolog.Logger
	e zerolog.Logger
}

// logger initialize.
// Open or create log files.
func init() {
	l, _ := NewFileLogger()
	infra.Logger = l
}

// NewFileLogger returns new fileLogger.
// fileLogger is implements of infra.ILogger.
func NewFileLogger() (*fileLogger, error) {
	ifp, err := open(env.InfoLogFile)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	wfp, err := open(env.WarnLogFile)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	efp, err := open(env.ErrLogFile)
	if err != nil {
		return nil, errors.New(err.Error())
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

// open is open the file placed in log locate and create it if it's not here.
// It is a wrapper for os.OpenFile, see there for details.
func open(file string) (*os.File, error) {
	path := filepath.Join(logLocate(), file)
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
}

// logLocate returns a string of the log file path.
func logLocate() string {
	return filepath.Join(util.RootDir(), env.LogDir)
}

var _ infra.ILogger = (*fileLogger)(nil)
