package middleware

import "net/http"

// Logger logs request method, path, and response status.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		_ = responseWriter
		_ = request
		// log request start
		// call next
		// log response status + duration
	}
}
