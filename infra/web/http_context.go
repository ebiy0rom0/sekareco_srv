package web

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sekareco_srv/domain/infra"
	infra_ "sekareco_srv/interface/infra"

	"github.com/gorilla/mux"
)

// A HttpContext is http.ResponseWriter and http.Request wrapper.
// Provides only the processing used in this application.
type HttpContext struct {
	Writer  http.ResponseWriter
	Request *http.Request
}

// NewHttpContext returns a new HttpContext wrapped the arguments w and r.
func NewHttpContext(w http.ResponseWriter, r *http.Request) *HttpContext {
	return &HttpContext{
		Writer:  w,
		Request: r,
	}
}

// Vars returns the uri query parameter map for route path.
//
// Returns map format is
// key: Specified string in the route path parameter.
// value: Entered fro uri parameter
func (c *HttpContext) Vars() map[string]string {
	return mux.Vars(c.Request)

}

// Decode reads the request body and stores it in the value pointed to by i.
func (c *HttpContext) Decode(i interface{}) error {
	err := json.NewDecoder(c.Request.Body).Decode(&i)
	if err != nil {
		return fmt.Errorf("bad request: %s", err)
	}
	return nil
}

// Response makes the response header and parameter body
// based on the arguments status code and v encoded in json.
//
// Argument v is desire be a key-value map
// because will encoding in json using json.Marshal().
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

// MakeError returns the HttpError converted into a format
// that can be passed to the Response argument v.
// Output to the error log is made at the same time.
func (c *HttpContext) MakeError(err error) *infra.HttpError {
	// TODO: output to the error log
	return &infra.HttpError{
		Error: err.Error(),
	}
}

// interface implementation checks
var _ infra_.HttpContext = (*HttpContext)(nil)
