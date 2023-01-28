//go:build !production && !test
// +build !production,!test

package env

const (
	LogDir      = "log/"
	ErrLogFile  = "error.log"
	WarnLogFile = "warn.log"
	InfoLogFile = "info.log"

	DbDir  = "db/"
	DbFile = "sekareco.db"
)
