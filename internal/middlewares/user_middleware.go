package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jim-nnamdi/kotts/internal/user"
)

var (
	jwt_key = []byte("kotts_secret_key")
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		token_string := c.Value
		claims := &user.DataToEncode{}
		token_gen, err := jwt.ParseWithClaims(token_string, claims, func(t *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token_gen.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Email)))
	})
}

func RefreshMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("user_token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token_gen := c.Value
		claims := &user.DataToEncode{}
		token_strip, err := jwt.ParseWithClaims(token_gen, claims, func(t *jwt.Token) (interface{}, error) {
			return jwt_key, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !token_strip.Valid {
			w.WriteHeader(http.StatusUnauthorized)
		}

		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 30 seconds of expiry. Otherwise, return a bad request status
		if time.Duration(claims.ExpiresAt) > time.Until(time.Now().Add(30*time.Second)) {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Now, create a new token for the current use, with a renewed expiration time
		expirationTime := time.Now().Add(5 * time.Minute)
		claims.ExpiresAt = expirationTime.Unix()
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token_gen, err = token.SignedString(jwt_key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "user_token",
			Value:   token_gen,
			Expires: expirationTime,
		})
	})
}
