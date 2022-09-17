package infra

type ILogger interface {
	Error(error)
	Warn(error)
	Info(error)
}

// Logger instance to be called from code.
var Logger ILogger
