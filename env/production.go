//go:build production
// +build production

package env

const (
	LogDir      = "log/"
	ErrLogFile  = "error.log"
	WarnLogFile = "warn.log"
	InfoLogFile = "info.log"

	DbDir  = "db/"
	DbFile = "sekareco.db"
)
