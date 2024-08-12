package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"admin": "password123",
}

var secretKey = []byte("hello")

func AuthenticationMiddlware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, req *http.Request) {

		var creds struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(req.Body).Decode(&creds); err != nil {
			http.Error(w, "Bad Request: Invalid JSON", http.StatusBadRequest)
			return
		}

		if pwd, exist := users[creds.Username]; !exist || pwd != creds.Password {
			http.Error(w, "Unauthorized: Invalid credentials", http.StatusUnauthorized)
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": creds.Username,
			"exp":      time.Now().Add(time.Minute * 30).Unix(),
		})

		tokenString, err := token.SignedString(secretKey)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Authorization", "Bearer "+tokenString)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Authentication successful, token generated"))

		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}
