package repository

import (
	"example.com/field-observation-ledger/internal/domain"
)

type Snapshot struct {
	Observations []domain.Observation
	Species      []domain.Species
	Sites        []domain.Site
}

func (c Catalog) Snapshot() Snapshot {
	return Snapshot{
		Observations: c.Observations.List(ObservationFilter{}),
		Species:      c.Species.List(),
		Sites:        c.Sites.List(),
	}
}

func (s Snapshot) ObservationCount() int {
	return len(s.Observations)
}

func (s Snapshot) CatalogCount() int {
	return len(s.Species) + len(s.Sites)
}
