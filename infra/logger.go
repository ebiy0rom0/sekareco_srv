package infra

import (
	"context"
	"io"
	"os"
	"sekareco_srv/domain/infra"

	"github.com/ebiy0rom0/errors"

	"github.com/rs/zerolog"
)

type logger struct {
	l zerolog.Logger
}

var loggerKey = struct{}{}

type logDetail struct {
	level infra.LogLevel
	msg   string
}

// logger initialize.
// Open or create log files.
func init() {
	newLogger(os.Stdout)
}

// NewFileLogger returns new fileLogger.
// fileLogger is implements of infra.ILogger.
func newLogger(w io.Writer) {
	infra.Logger = &logger{
		l: zerolog.New(w),
	}
}

func NewLogDetail(ctx context.Context) context.Context {
	return context.WithValue(ctx, &loggerKey, &logDetail{})
}

func getLogDetail(ctx context.Context) (*logDetail, error) {
	v := ctx.Value(&loggerKey)
	detail, ok := v.(*logDetail)
	if !ok {
		return nil, errors.New("log detail not set in context")
	}
	return detail, nil
}

func (l *logger) Log(ctx context.Context, msg string, level infra.LogLevel) error {
	detail, err := getLogDetail(ctx)
	if err != nil {
		return errors.WithStack(err)
	}

	if detail.level < level {
		detail.level = level
	}
	if len(detail.msg) == 0 {
		detail.msg = msg
	}
	return nil
}

// Error writes info message into info log file.
func (l *logger) Info(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, infra.LogLevelInfo)
}

// Error writes warning message into warn log file.
func (l *logger) Warn(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, infra.LogLevelWarn)
}

// Error writes error message into error log file.
func (l *logger) Err(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, infra.LogLevelErr)
}

// Error writes error message into error log file.
func (l *logger) Fatal(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, infra.LogLevelErr)
}

// Error writes error message into error log file.
func (l *logger) Debug(ctx context.Context, msg string) error {
	return l.Log(ctx, msg, infra.LogLevelErr)
}

func (l *logger) Send(ctx context.Context) error {
	detail, err := getLogDetail(ctx)
	if err != nil {
		return err
	}

	switch detail.level {
	case infra.LogLevelInfo:
		l.l.Info().Str("message", detail.msg).Send()
	case infra.LogLevelWarn:
		l.l.Warn().Str("message", detail.msg).Send()
	case infra.LogLevelErr:
		l.l.Error().Str("message", detail.msg).Send()
	case infra.LogLevelFatal:
		l.l.Fatal().Str("message", detail.msg).Send()
	case infra.LogLevelDebug:
		l.l.Debug().Str("message", detail.msg).Send()
	}
	return nil
}

var _ infra.ILogger = (*logger)(nil)
