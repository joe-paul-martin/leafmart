package middleware

import "net/http"

func AuthenticationMiddlware(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, req *http.Request) {
		next.ServeHTTP(w, req)
	}

	return http.HandlerFunc(fn)
}
