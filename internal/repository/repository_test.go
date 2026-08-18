package repository

import (
	"testing"
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

func TestObservationRepositoryFiltersAndSorts(t *testing.T) {
	repo := NewObservationRepository()
	older := domain.Observation{ID: "old", SiteID: "site-a", Status: domain.StatusDraft, ObservedAt: time.Unix(1, 0)}
	newer := domain.Observation{ID: "new", SiteID: "site-a", Status: domain.StatusApproved, ObservedAt: time.Unix(2, 0)}
	if err := repo.Create(newer); err != nil {
		t.Fatal(err)
	}
	if err := repo.Create(older); err != nil {
		t.Fatal(err)
	}
	items := repo.List(ObservationFilter{SiteID: "site-a"})
	if len(items) != 2 || items[0].ID != "old" {
		t.Fatalf("unexpected filtered order: %#v", items)
	}
}

func TestEvidenceVerification(t *testing.T) {
	repo := NewEvidenceRepository()
	item := domain.Evidence{
		ID: "ev-1", ObservationID: "obs-1", Reference: "field/photo-1", Observer: "tester",
	}
	if err := repo.Create(item); err != nil {
		t.Fatal(err)
	}
	verified, err := repo.Verify("ev-1")
	if err != nil || !verified.Verified {
		t.Fatalf("expected verified evidence, got %#v, %v", verified, err)
	}
}
