package analysis

import (
	"testing"
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

func TestRankSpeciesAndWindow(t *testing.T) {
	items := []domain.Observation{
		{ID: "a", SiteID: "site-a", SpeciesID: "sp-a", Count: 2, ObservedAt: time.Date(2026, 8, 1, 0, 0, 0, 0, time.UTC)},
		{ID: "b", SiteID: "site-a", SpeciesID: "sp-b", Count: 4, ObservedAt: time.Date(2026, 8, 3, 0, 0, 0, 0, time.UTC)},
	}
	ranks := RankSpecies(items, 5)
	if len(ranks) != 2 || ranks[0].SpeciesID != "sp-b" || ranks[0].Position != 1 {
		t.Fatalf("unexpected ranks: %#v", ranks)
	}
	window := ObservationWindow(items)
	if window.Days != 3 {
		t.Fatalf("expected three-day window, got %#v", window)
	}
}

func TestQualityAndCoverage(t *testing.T) {
	items := []domain.Observation{
		{ID: "a", SiteID: "site-a", SpeciesID: "sp-a", Count: 1, ObservedAt: time.Now()},
		{ID: "b", SiteID: "site-b", SpeciesID: "sp-a", Count: 1},
	}
	report := Quality(items)
	if report.CompleteRecords != 1 || report.IncompleteRecords != 1 {
		t.Fatalf("unexpected quality report: %#v", report)
	}
	coverage := BuildCoverage(items)
	if coverage.Sites != 2 || coverage.Species != 1 {
		t.Fatalf("unexpected coverage: %#v", coverage)
	}
}
