package infra

type HttpContext interface {
	Vars() map[string]string
	Decode(interface{}) error
	Response(int, ...interface{}) error
	MakeError(error) map[string]string
}
