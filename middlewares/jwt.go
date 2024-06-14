package middlewares

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-dating-app/helpers"
	"net/http"
	"strings"
)

// use env file for security but in this test just use string
var JwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func JwtAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helpers.Response{
				Message: "Authorization header missing",
			})
			return
		}

		const bearerPrefix = "Bearer "
		if !strings.HasPrefix(authHeader, bearerPrefix) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helpers.Response{
				Message: "Invalid authorization header format",
			})
			return
		}

		tokenString := authHeader[len(bearerPrefix):]

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(helpers.Response{
				Message: "Invalid token",
			})
			return
		}

		// Pass the request to the next handler
		next(w, r)
	}
}
