package logger

import "sekareco_srv/domain/infra"

type awsLogger struct {
	// needs objects
}

func NewAwsLogger() (*awsLogger, error) {
	return &awsLogger{}, nil
}

func (l *awsLogger) Info (err error) {
}

func (l *awsLogger) Warn (err error) {
}

func (l *awsLogger) Error (err error) {
}

var _ infra.ILogger = (*awsLogger)(nil)