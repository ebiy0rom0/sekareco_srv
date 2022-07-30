package infra

import "sekareco_srv/domain/infra"

type HttpContext interface {
	Vars() map[string]string
	Decode(interface{}) error
	Response(int, interface{}) error
	MakeError(error) *infra.HttpError
}
