package logger

import (
	"sekareco_srv/domain/infra"
)

// InitLogger initializes infra.Logger.
// Switch logger type by stage.
func InitLogger(stage string) error {
	var l infra.ILogger
	var err error

	switch stage {
	case "prod":
		if l, err = NewAwsLogger(); err != nil {
			return err
		}
	case "dev":
		if l, err = NewFileLogger(); err != nil {
			return err
		}
	default:
		if l, err = NewFileLogger(); err != nil {
			return err
		}
	}

	infra.Logger = l
	return nil
}
