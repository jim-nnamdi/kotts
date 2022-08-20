package middlewares

import (
	"log"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusOK)
				return
			}
		}()
		next.ServeHTTP(w, r)
	})
}
