package infra

import "context"

type LogLevel int

const (
	LogLevelInfo LogLevel = iota + 1
	LogLevelWarn
	LogLevelErr
	LogLevelFatal
	LogLevelDebug
)

type ILogger interface {
	Log(context.Context, string, LogLevel) error
	Info(context.Context, string) error
	Warn(context.Context, string) error
	Err(context.Context, string) error
	Fatal(context.Context, string) error
	Debug(context.Context, string) error
	Send(context.Context) error
}

// Logger instance to be called from code.
var Logger ILogger
