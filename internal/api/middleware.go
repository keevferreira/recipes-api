package api

import (
	"log"
	"net/http"
	"time"
)

// MiddlewareLogging é um middleware para registrar informações sobre as requisições HTTP.
func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Delega o controle para o próximo handler
		next.ServeHTTP(w, r)

		// Registra informações sobre a requisição
		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
			time.Since(start),
		)
	})
}
