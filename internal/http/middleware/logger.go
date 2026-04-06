package middleware

import "net/http"

// Logger logs request method, path, and response status.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// log request start
		// call next
		// log response status + duration
	}
}
