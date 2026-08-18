package validate

import (
	"errors"
	"strings"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrMissingName = errors.New("name is required")
var ErrMissingScientificName = errors.New("scientific name is required")
var ErrMissingRegion = errors.New("region is required")

func Species(item domain.Species) error {
	if strings.TrimSpace(item.ID) == "" || strings.TrimSpace(item.CommonName) == "" {
		return ErrMissingName
	}
	if strings.TrimSpace(item.ScientificName) == "" {
		return ErrMissingScientificName
	}
	return nil
}

func Site(item domain.Site) error {
	if strings.TrimSpace(item.ID) == "" || strings.TrimSpace(item.Name) == "" {
		return ErrMissingName
	}
	if strings.TrimSpace(item.Region) == "" {
		return ErrMissingRegion
	}
	if item.Latitude < -90 || item.Latitude > 90 {
		return errors.New("latitude is out of range")
	}
	if item.Longitude < -180 || item.Longitude > 180 {
		return errors.New("longitude is out of range")
	}
	return nil
}
