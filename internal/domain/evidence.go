package domain

import "time"

type EvidenceKind string

const (
	EvidencePhoto EvidenceKind = "photo"
	EvidenceAudio EvidenceKind = "audio"
	EvidenceField EvidenceKind = "field_note"
)

type Evidence struct {
	ID            string       `json:"id"`
	ObservationID string       `json:"observation_id"`
	Kind          EvidenceKind `json:"kind"`
	Reference     string       `json:"reference"`
	CapturedAt    time.Time    `json:"captured_at"`
	Observer      string       `json:"observer"`
	Verified      bool         `json:"verified"`
}

func (e Evidence) IsSupported() bool {
	return e.Kind == EvidencePhoto || e.Kind == EvidenceAudio || e.Kind == EvidenceField
}

func (e Evidence) IsComplete() bool {
	return e.ID != "" && e.ObservationID != "" && e.Reference != "" && !e.CapturedAt.IsZero()
}

func (e Evidence) IsReviewReady() bool {
	return e.IsSupported() && e.IsComplete() && e.Verified
}
