package repository

import (
    "testing"
    "time"

    "example.com/field-observation-ledger/internal/domain"
)

func TestNoteTagsAreIsolated(t *testing.T) {
    repo := NewNoteRepository()
    err := repo.Create(domain.FieldNote{ID: "note-1", ObservationID: "obs-1", Author: "qa", Text: "ridge", Tags: []string{"habitat"}, CreatedAt: time.Unix(1, 0)})
    if err != nil {
        t.Fatal(err)
    }
    listed := repo.ListForObservation("obs-1")
    listed[0].Tags[0] = "mutated"
    listedAgain := repo.ListForObservation("obs-1")
    if listedAgain[0].Tags[0] != "habitat" {
        t.Fatalf("expected stored tag to remain habitat, got %#v", listedAgain[0].Tags)
    }
}
