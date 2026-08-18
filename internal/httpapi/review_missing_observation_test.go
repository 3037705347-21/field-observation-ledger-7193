package httpapi_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"example.com/field-observation-ledger/internal/app"
	"example.com/field-observation-ledger/internal/config"
)

func TestMissingObservationReturnsNotFound(t *testing.T) {
	handler := app.New(config.Settings{}).Handler()
	request := httptest.NewRequest(http.MethodPost, "/api/v1/reviews/obs-missing/approve", strings.NewReader(`{"reviewer":"qa"}`))
	request.Header.Set("Content-Type", "application/json")
	result := httptest.NewRecorder()
	handler.ServeHTTP(result, request)

	if result.Code != http.StatusNotFound {
		t.Fatalf("expected missing observation to return %d, got %d: %s", http.StatusNotFound, result.Code, result.Body.String())
	}
}
