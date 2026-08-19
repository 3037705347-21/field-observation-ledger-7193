package httpapi_test

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"

    "example.com/field-observation-ledger/internal/app"
    "example.com/field-observation-ledger/internal/config"
)

func TestObservationListZeroLimitUsesDefault(t *testing.T) {
    handler := app.New(config.Settings{}).Handler()
    req := httptest.NewRequest(http.MethodGet, "/api/v1/observations?limit=0", nil)
    result := httptest.NewRecorder()
    handler.ServeHTTP(result, req)
    if result.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d: %s", result.Code, result.Body.String())
    }
    body := result.Body.String()
    if !strings.Contains(body, `"id":"obs-001"`) || !strings.Contains(body, `"id":"obs-002"`) {
        t.Fatalf("expected default page to include seeded observations, got %s", body)
    }
    if !strings.Contains(body, `"total":2`) {
        t.Fatalf("expected total 2, got %s", body)
    }
}
