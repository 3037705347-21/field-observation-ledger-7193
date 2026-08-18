package service

import (
	"testing"
	"time"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/metrics"
	"example.com/field-observation-ledger/internal/repository"
)

func TestSummaryReviewEventsReflectHistory(t *testing.T) {
	catalog := repository.NewCatalog()
	timeSource := fixedClock{value: time.Date(2026, 8, 18, 9, 0, 0, 0, time.UTC)}
	observations := NewObservationService(catalog.Observations, catalog.Sites, catalog.Species, timeSource, metrics.NewCounter())
	item, err := observations.Create(domain.Observation{
		ID: "obs-summary-review", SiteID: "site-alpine", SpeciesID: "sp-robin",
		ObservedAt: time.Date(2026, 8, 18, 8, 0, 0, 0, time.UTC), Count: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	reviews := NewReviewService(observations, catalog.Reviews, timeSource)
	if _, _, err := reviews.Decide(item.ID, "reviewer", domain.DecisionApprove, "accepted"); err != nil {
		t.Fatal(err)
	}

	summary := NewSummaryService(catalog.Observations, catalog.Reviews).Build(repository.ObservationFilter{})
	if summary.Analysis.ReviewEvents != 1 {
		t.Fatalf("expected one review event from review history, got %d", summary.Analysis.ReviewEvents)
	}
}
