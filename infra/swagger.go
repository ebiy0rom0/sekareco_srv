package infra

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
)

// for debug
func SwaggerUI() http.Handler {
	return httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/swagger/doc.json"),
	)
}
