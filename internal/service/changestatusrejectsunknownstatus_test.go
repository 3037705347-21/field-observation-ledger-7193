package service

import (
    "testing"

    "example.com/field-observation-ledger/internal/domain"
    "example.com/field-observation-ledger/internal/metrics"
    "example.com/field-observation-ledger/internal/repository"
)

func TestChangeStatusRejectsUnknownStatus(t *testing.T) {
    catalog := repository.NewCatalog()
    observations := NewObservationService(catalog.Observations, catalog.Sites, catalog.Species, fixedClock{}, metrics.NewCounter())
    updated, err := observations.ChangeStatus("obs-002", domain.ObservationStatus("archived"))
    if err == nil {
        t.Fatalf("expected unknown status to be rejected, got %#v", updated)
    }
    if updated.Status != "" {
        t.Fatalf("expected no updated item, got %#v", updated)
    }
}
