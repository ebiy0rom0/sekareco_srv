package handler

type HttpContext interface {
	Vars() map[string]string
	Decode(interface{}) error
	Response(int, interface{}) error
	MakeError(string) map[string]string
}
