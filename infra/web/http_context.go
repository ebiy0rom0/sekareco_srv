package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sekareco_srv/domain/infra"

	"github.com/gorilla/mux"
)

type HttpContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

func NewHttpContext(w http.ResponseWriter, r *http.Request) *HttpContext {
	return &HttpContext{
		Writer:  w,
		Request: r,
	}
}

func (c *HttpContext) Vars() map[string]string {
	return mux.Vars(c.Request)
}

func (c *HttpContext) Decode(i interface{}) error {
	err := json.NewDecoder(c.Request.Body).Decode(&i)
	if err != nil {
		return fmt.Errorf("bad request: %s", err)
	}
	return nil
}

func (c *HttpContext) Response(code int, v interface{}) error {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)

	if v != nil {
		output, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("marshal failed: %s", err)
		}
		c.Writer.Write(output)
	}

	return nil
}

func (c *HttpContext) MakeError(err error) *infra.HttpError {
	return &infra.HttpError{
		Error: err.Error(),
	}
}
