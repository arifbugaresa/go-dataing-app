package user

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"go-dating-app/databases"
	"go-dating-app/helpers"
	"go-dating-app/middlewares"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user UserReq
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(helpers.Response{
			Message: "Failed Decode",
			Data:    nil,
		})
		return
	}

	var dbUser UserReq
	if err := databases.DBConnection.Where("username = ? AND password = ?", user.Username, user.Password).First(&dbUser).Error; err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(helpers.Response{
			Message: "Invalid username or password",
			Data:    nil,
		})
		return
	}

	// Set expiration time for the token
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims, which includes the username and expiry time
	claims := &middlewares.Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the specified algorithm and claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(middlewares.JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(helpers.Response{
			Message: "Internal Server Error",
		})
		return
	}

	// Return the token
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(helpers.Response{
		Message: "Login successful",
		Data:    map[string]string{"token": tokenString},
	})
}
