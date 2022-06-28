package middleware

import (
	"net/http"

	"github.com/rs/cors"
)

func InitCors() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:8080",
		},
		AllowedHeaders: []string{
			"Content-Type",
		},
		AllowedMethods: []string{
			http.MethodHead, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions,
		},
		AllowCredentials: true,
	})
}
