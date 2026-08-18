package service

import (
	"testing"
	"time"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/metrics"
	"example.com/field-observation-ledger/internal/repository"
)

type fixedClock struct {
	value time.Time
}

func (c fixedClock) Now() time.Time {
	return c.value
}

func TestObservationServiceCreatesDraftAndCountsMetric(t *testing.T) {
	catalog := repository.NewCatalog()
	counter := metrics.NewCounter()
	service := NewObservationService(
		catalog.Observations, catalog.Sites, catalog.Species,
		fixedClock{value: time.Date(2026, 8, 18, 9, 0, 0, 0, time.UTC)}, counter,
	)
	item, err := service.Create(domain.Observation{
		ID: "obs-test", SiteID: "site-alpine", SpeciesID: "sp-robin",
		ObservedAt: time.Now(), Count: 2,
	})
	if err != nil {
		t.Fatal(err)
	}
	if item.Status != domain.StatusDraft || counter.Get("observations.created") != 1 {
		t.Fatalf("unexpected creation result: %#v, %d", item, counter.Get("observations.created"))
	}
}

func TestReviewServiceChangesState(t *testing.T) {
	catalog := repository.NewCatalog()
	timeSource := fixedClock{value: time.Date(2026, 8, 18, 9, 0, 0, 0, time.UTC)}
	observations := NewObservationService(catalog.Observations, catalog.Sites, catalog.Species, timeSource, metrics.NewCounter())
	item, err := observations.Create(domain.Observation{
		ID: "obs-review", SiteID: "site-alpine", SpeciesID: "sp-robin", ObservedAt: time.Now(), Count: 1,
	})
	if err != nil {
		t.Fatal(err)
	}
	reviews := NewReviewService(observations, catalog.Reviews, timeSource)
	updated, review, err := reviews.Decide(item.ID, "reviewer", domain.DecisionApprove, "accepted")
	if err != nil || updated.Status != domain.StatusApproved || review.Decision != domain.DecisionApprove {
		t.Fatalf("unexpected review result: %#v, %#v, %v", updated, review, err)
	}
}
