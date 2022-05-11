package web

import (
	"encoding/json"
	"fmt"
	"net/http"

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

// TODO: bag - unmarshal failed
func (c *HttpContext) Decode(i interface{}) error {
	err := json.NewDecoder(c.Request.Body).Decode(&i)
	if err != nil {
		return fmt.Errorf("bad request: %s", err)
	}
	return nil
}

func (c *HttpContext) Response(code int, v interface{}) error {
	output, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("marshal failed: %s", err)
	}

	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(code)
	c.Writer.Write(output)

	return nil
}
