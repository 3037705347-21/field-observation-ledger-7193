package repository

import (
	"testing"
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

func TestNoteCreateCopiesTags(t *testing.T) {
	repo := NewNoteRepository()
	tags := []string{"habitat", "ridge"}
	item := domain.FieldNote{
		ID:            "note-create",
		ObservationID: "obs-1",
		Author:        "qa",
		Text:          "ridge note",
		Tags:          tags,
		CreatedAt:     time.Unix(2, 0),
	}
	if err := repo.Create(item); err != nil {
		t.Fatal(err)
	}
	tags[0] = "mutated"

	got := repo.ListForObservation("obs-1")
	if len(got) != 1 || got[0].Tags[0] != "habitat" {
		t.Fatalf("expected stored tags to remain unchanged, got %#v", got)
	}
}

func TestNoteListCopiesTagsAndPreservesFields(t *testing.T) {
	repo := NewNoteRepository()
	older := domain.FieldNote{
		ID:            "note-older",
		ObservationID: "obs-1",
		Author:        "first",
		Text:          "low trail",
		Tags:          []string{"habitat"},
		CreatedAt:     time.Unix(1, 0),
	}
	newer := domain.FieldNote{
		ID:            "note-newer",
		ObservationID: "obs-1",
		Author:        "second",
		Text:          "high ridge",
		Tags:          []string{"elevation"},
		CreatedAt:     time.Unix(2, 0),
	}
	if err := repo.Create(newer); err != nil {
		t.Fatal(err)
	}
	if err := repo.Create(older); err != nil {
		t.Fatal(err)
	}

	listed := repo.ListForObservation("obs-1")
	if len(listed) != 2 {
		t.Fatalf("expected two notes, got %#v", listed)
	}
	listed[0].Tags[0] = "mutated"
	if listed[0].ID != "note-older" || listed[0].Author != "first" || listed[0].Text != "low trail" {
		t.Fatalf("expected oldest note fields to be preserved, got %#v", listed[0])
	}
	if listed[1].ID != "note-newer" || listed[1].CreatedAt != newer.CreatedAt {
		t.Fatalf("expected notes sorted by creation time, got %#v", listed)
	}

	again := repo.ListForObservation("obs-1")
	if again[0].Tags[0] != "habitat" || again[1].Tags[0] != "elevation" {
		t.Fatalf("expected stored tags to remain unchanged, got %#v", again)
	}
}
