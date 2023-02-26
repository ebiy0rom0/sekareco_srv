package infra

type HttpError struct {
	Code int
	Msg  string
}

type HttpContext interface {
	Vars() map[string]string
	Decode(...interface{}) error
	Response(int, interface{}) error
}
