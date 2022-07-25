package infra

type HttpError struct {
	Error string `json:"error" example:"set a server error message"`
}
