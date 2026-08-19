package domain

import "testing"

func TestFieldNoteHasTagMatchesNormalizedTag(t *testing.T) {
    note := FieldNote{ID: "note-1", ObservationID: "obs-1", Author: "qa", Text: "ridge", Tags: []string{"habitat"}}
    if !note.HasTag("Habitat") {
        t.Fatalf("expected case-insensitive tag match")
    }
}
