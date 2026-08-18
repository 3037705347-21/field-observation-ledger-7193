package httpapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/field-observation-ledger/internal/app"
	"example.com/field-observation-ledger/internal/config"
)

func TestHealthAndCreateObservation(t *testing.T) {
	handler := app.New(config.Settings{}).Handler()
	health := httptest.NewRecorder()
	handler.ServeHTTP(health, httptest.NewRequest(http.MethodGet, "/health", nil))
	if health.Code != http.StatusOK || !strings.Contains(health.Body.String(), `"status":"ok"`) {
		t.Fatalf("unexpected health response: %d %s", health.Code, health.Body.String())
	}
	request := httptest.NewRequest(http.MethodPost, "/api/v1/observations", strings.NewReader(
		`{"id":"obs-http","site_id":"site-alpine","species_id":"sp-robin","observed_at":"2026-08-18T09:30:00Z","count":2}`,
	))
	request.Header.Set("Content-Type", "application/json")
	result := httptest.NewRecorder()
	handler.ServeHTTP(result, request)
	if result.Code != http.StatusCreated || !strings.Contains(result.Body.String(), `"status":"draft"`) {
		t.Fatalf("unexpected create response: %d %s", result.Code, result.Body.String())
	}
}
