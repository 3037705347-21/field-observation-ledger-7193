package analysis

import (
	"testing"
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

func TestObservationWindowUsesCalendarDays(t *testing.T) {
	items := []domain.Observation{
		{ID: "first", ObservedAt: time.Date(2026, 8, 1, 23, 30, 0, 0, time.UTC)},
		{ID: "last", ObservedAt: time.Date(2026, 8, 2, 0, 30, 0, 0, time.UTC)},
	}

	window := ObservationWindow(items)
	if window.Days != 2 {
		t.Fatalf("expected two calendar days, got %d", window.Days)
	}
}
