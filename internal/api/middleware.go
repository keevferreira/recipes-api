package api

import (
	"log"
	"net/http"
)

// MiddlewareLogging é um middleware para registrar informações sobre as requisições HTTP.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Logging
		log.Printf("Request: %s %s", r.Method, r.RequestURI)

		// Chamada para o próximo handler
		next.ServeHTTP(w, r)
	})
}
