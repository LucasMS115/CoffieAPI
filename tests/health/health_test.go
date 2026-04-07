package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	apphttp "coffie/internal/http"
)

func TestHealthEndpoint(testingContext *testing.T) {
	// given
	server := apphttp.NewServer(":8080", nil)

	// when
	request := httptest.NewRequest("GET", "/health", nil)
	responseRecorder := httptest.NewRecorder()
	server.Handler.ServeHTTP(responseRecorder, request)

	// then
	if responseRecorder.Code != http.StatusOK {
		testingContext.Errorf("expected status %d, got %d", http.StatusOK, responseRecorder.Code)
	}

	var responseBody map[string]string
	if decodeError := json.NewDecoder(responseRecorder.Body).Decode(&responseBody); decodeError != nil {
		testingContext.Fatalf("expected valid JSON, got error: %v", decodeError)
	}

	if responseBody["status"] != "ok" {
		testingContext.Errorf("expected status 'ok', got %q", responseBody["status"])
	}
}
