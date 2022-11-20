package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

// NewCorsConfig returns a new Cors handler.
func NewCorsConfig() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",
		},
		AllowedHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Authorization",
		},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowCredentials: false,
		Debug:            true,
	})
}

// NewCorsConfigForMainte returns a new Cors handler used by bastion server.
func NewCorsConfigForMainte() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"*", // TODO: bastion server's domain
		},
		AllowedHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Authorization",
		},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowCredentials: false,
		Debug:            true,
	})
}
