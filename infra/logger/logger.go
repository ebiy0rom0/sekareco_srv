package logger

import (
	"sekareco_srv/domain/infra"

	"github.com/ebiy0rom0/errors"
)

// InitLogger initializes infra.Logger.
// Switch logger type by stage.
func InitLogger(stage string) error {
	var l infra.ILogger
	var err error

	switch stage {
	case "prod":
		l, err = NewAwsLogger()
		if err != nil {
			return errors.New(err.Error())
		}
	case "dev":
		l, err = NewFileLogger()
		if err != nil {
			return errors.New(err.Error())
		}
	default:
		l, err = NewFileLogger()
		if err != nil {
			return errors.New(err.Error())
		}
	}

	infra.Logger = l
	return nil
}
