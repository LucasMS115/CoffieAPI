package integration

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	apphttp "coffie/internal/http"
)

func TestHealthEndpoint(t *testing.T) {
	// given
	server := apphttp.NewServer(":8080")

	// when
	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()
	server.Handler.ServeHTTP(rec, req)

	// then
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rec.Code)
	}

	var body map[string]string
	if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
		t.Fatalf("expected valid JSON, got error: %v", err)
	}

	if body["status"] != "ok" {
		t.Errorf("expected status 'ok', got %q", body["status"])
	}
}
