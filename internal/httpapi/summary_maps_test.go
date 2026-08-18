package httpapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/field-observation-ledger/internal/app"
	"example.com/field-observation-ledger/internal/config"
)

func TestSummaryReturnsStableMaps(t *testing.T) {
	handler := app.New(config.Settings{}).Handler()
	request := httptest.NewRequest(http.MethodGet, "/api/v1/summary", nil)
	result := httptest.NewRecorder()
	handler.ServeHTTP(result, request)

	if result.Code != http.StatusOK {
		t.Fatalf("expected summary status %d, got %d: %s", http.StatusOK, result.Code, result.Body.String())
	}
	body := result.Body.String()
	if strings.Contains(body, `"by_species":null`) || strings.Contains(body, `"by_status":null`) {
		t.Fatalf("summary maps must be JSON objects, got %s", body)
	}
}
