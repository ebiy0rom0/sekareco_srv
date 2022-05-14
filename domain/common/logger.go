package common

type Logger interface {
	Error(error)
	Warn(error)
	Info(error)
}
