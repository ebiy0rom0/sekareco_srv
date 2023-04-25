package infra

import (
	"fmt"
	"io"
	"os"
	"sekareco_srv/domain/infra"
	appLogger "sekareco_srv/logger"

	"github.com/rs/zerolog"
)

type logger struct {
	l zerolog.Logger
}

func init() {
	// logger are automatically initialized with os.Stdout.
	// When changing the output distination, call InitLogger() to reinitialize it.
	appLogger.L = &logger{l: zerolog.New(os.Stdout)}
}

// InitLogger initialize logger.
func InitLogger(w io.Writer) {
	appLogger.L = &logger{l: zerolog.New(w)}
}

// Log writes a message at the specified level to the configured os.Writer.
func (l *logger) Log(level infra.LogLevel, message string) {
	switch level {
	case infra.LogLevelInfo:
		l.l.Info().Msg(message)
	case infra.LogLevelWarn:
		l.l.Warn().Msg(message)
	case infra.LogLevelErr:
		l.l.Error().Msg(message)
	case infra.LogLevelFatal:
		// zerolog.Logger.Fatal() executes os.Exit(1).
		// It's not used because it terminates the program.
		// Output with Error level and identify by appending level.
		l.l.Error().Str("level", "FATAL").Msg(message)
	case infra.LogLevelDebug:
		l.l.Debug().Msg(message)
	default:
		// Identified by appending level as well as fatal.
		l.l.Error().Str("level", "undeclared").Msg(message)
	}
}

// Logf writes a message to the configured os.Writer.
// You can use the format of the fmt package.
func (l *logger) Logf(level infra.LogLevel, format string, args ...any) {
	message := fmt.Sprintf(format, args...)
	l.Log(level, message)
}

// Info writes a message to the configured os.Writer.
func (l *logger) Info(message string) {
	l.Log(infra.LogLevelInfo, message)
}

// Infof writes a message to the configured os.Writer.
// You can use the format of the fmt package.
func (l *logger) Infof(format string, args ...any) {
	l.Logf(infra.LogLevelInfo, format, args...)
}

// Warn writes a message to the configured os.Writer.
func (l *logger) Warn(message string) {
	l.Log(infra.LogLevelWarn, message)
}

// Warnf writes a message to the configured os.Writer.
// You can use the format of the fmt package.
func (l *logger) Warnf(format string, args ...any) {
	l.Logf(infra.LogLevelInfo, format, args...)
}

// Error writes a message to the configured os.Writer.
func (l *logger) Error(message string) {
	l.Log(infra.LogLevelErr, message)
}

// Errorf writes a message to the configured os.Writer.
// You can use the format of the fmt package.
func (l *logger) Errorf(format string, args ...any) {
	l.Logf(infra.LogLevelErr, format, args...)
}

// Fatal writes a message to the configured os.Writer.
func (l *logger) Fatal(message string) {
	l.Log(infra.LogLevelFatal, message)
}

// Fatalf writes a message to the configured os.Writer.
// You can use the format of the fmt package.
func (l *logger) Fatalf(format string, args ...any) {
	l.Logf(infra.LogLevelFatal, format, args...)
}

// Debug writes a message to the configured os.Writer.
func (l *logger) Debug(message string) {
	l.Log(infra.LogLevelDebug, message)
}

// Debugf writes a message to the configured os.Writer.
// You can use the format of the fmt package.
func (l *logger) Debugf(format string, args ...any) {
	l.Logf(infra.LogLevelDebug, format, args...)
}

var _ infra.Logger = (*logger)(nil)
