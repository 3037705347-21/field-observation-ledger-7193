package domain

import "time"

type ObservationStatus string

const (
	StatusDraft    ObservationStatus = "draft"
	StatusApproved ObservationStatus = "approved"
	StatusRejected ObservationStatus = "rejected"
)

type Observation struct {
	ID         string            `json:"id"`
	SiteID     string            `json:"site_id"`
	SpeciesID  string            `json:"species_id"`
	ObservedAt time.Time         `json:"observed_at"`
	Count      int               `json:"count"`
	Notes      string            `json:"notes,omitempty"`
	Status     ObservationStatus `json:"status"`
	CreatedAt  time.Time         `json:"created_at"`
	UpdatedAt  time.Time         `json:"updated_at"`
}

func (o Observation) IsReviewable() bool {
	return o.Status == StatusDraft
}

func (o Observation) IsVisible() bool {
	return o.Status != ""
}
