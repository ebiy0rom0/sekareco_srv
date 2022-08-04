package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func InitCors() *cors.Cors {
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
