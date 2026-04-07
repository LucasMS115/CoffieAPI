package response

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON response with the given status code.
func JSON(responseWriter http.ResponseWriter, status int, data interface{}) {
	responseWriter.Header().Set("Content-Type", "application/json")
	responseWriter.WriteHeader(status)
	json.NewEncoder(responseWriter).Encode(data)
}

// ErrorResponse represents a standardized error payload.
type ErrorResponse struct {
	ErrorCode string       `json:"error_code"`
	Message   string       `json:"message"`
	Fields    []FieldError `json:"fields,omitempty"`
}

// FieldError describes a validation error on a specific field.
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Error writes a JSON error response.
func Error(responseWriter http.ResponseWriter, status int, code, message string, fields []FieldError) {
	JSON(responseWriter, status, ErrorResponse{
		ErrorCode: code,
		Message:   message,
		Fields:    fields,
	})
}
