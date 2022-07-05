package database

import "context"

type ExecFunc func(context.Context) (interface{}, error)

type SqlTransaction interface {
	Do(context.Context, ExecFunc) (interface{}, error)
}
