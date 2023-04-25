package infra

type LogLevel int

const (
	LogLevelInfo LogLevel = iota + 1
	LogLevelWarn
	LogLevelErr
	LogLevelFatal
	LogLevelDebug
)

type Logger interface {
	Log(level LogLevel, message string)
	Logf(level LogLevel, format string, args ...any)
	Info(message string)
	Infof(format string, args ...any)
	Warn(message string)
	Warnf(format string, args ...any)
	Error(message string)
	Errorf(format string, args ...any)
	Fatal(message string)
	Fatalf(format string, args ...any)
	Debug(message string)
	Debugf(format string, args ...any)
}
