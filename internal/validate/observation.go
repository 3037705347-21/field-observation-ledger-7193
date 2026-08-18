package validate

import (
	"errors"
	"strings"
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrMissingObservationID = errors.New("observation id is required")
var ErrMissingSiteID = errors.New("site id is required")
var ErrMissingSpeciesID = errors.New("species id is required")
var ErrInvalidCount = errors.New("count must be positive")
var ErrInvalidObservedAt = errors.New("observed_at must be RFC3339")

func Observation(item domain.Observation) error {
	if strings.TrimSpace(item.ID) == "" {
		return ErrMissingObservationID
	}
	if strings.TrimSpace(item.SiteID) == "" {
		return ErrMissingSiteID
	}
	if strings.TrimSpace(item.SpeciesID) == "" {
		return ErrMissingSpeciesID
	}
	if item.Count <= 0 {
		return ErrInvalidCount
	}
	if item.ObservedAt.IsZero() || item.ObservedAt.Location() == nil {
		return ErrInvalidObservedAt
	}
	return nil
}

func ParseObservedAt(value string) (time.Time, error) {
	return time.Parse(time.RFC3339, value)
}
