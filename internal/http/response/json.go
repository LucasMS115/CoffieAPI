package response

import (
	"encoding/json"
	"net/http"
)

// JSON writes a JSON response with the given status code.
func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
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
func Error(w http.ResponseWriter, status int, code, message string, fields []FieldError) {
	JSON(w, status, ErrorResponse{
		ErrorCode: code,
		Message:   message,
		Fields:    fields,
	})
}
