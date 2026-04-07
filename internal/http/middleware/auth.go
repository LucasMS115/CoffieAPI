package middleware

import "net/http"

// Auth validates JWT tokens and injects user_id into the request context.
func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(responseWriter http.ResponseWriter, request *http.Request) {
		_ = responseWriter
		_ = request
		// extract token from Authorization header
		// validate token
		// inject user_id into request context
		// call next
	}
}
