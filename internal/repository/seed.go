package repository

import (
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

type Catalog struct {
	Observations *ObservationRepository
	Species      *SpeciesRepository
	Sites        *SiteRepository
	Reviews      *ReviewRepository
	Evidence     *EvidenceRepository
	Notes        *NoteRepository
}

func NewCatalog() Catalog {
	catalog := Catalog{
		Observations: NewObservationRepository(),
		Species:      NewSpeciesRepository(),
		Sites:        NewSiteRepository(),
		Reviews:      NewReviewRepository(),
		Evidence:     NewEvidenceRepository(),
		Notes:        NewNoteRepository(),
	}
	seedCatalog(catalog)
	return catalog
}

func seedCatalog(catalog Catalog) {
	_ = catalog.Sites.Create(domain.Site{
		ID: "site-alpine", Name: "Alpine Ridge", Region: "North Range",
		Latitude: 45.112, Longitude: 7.431, Habitats: []string{"conifer", "meadow"},
	})
	_ = catalog.Sites.Create(domain.Site{
		ID: "site-marsh", Name: "Silver Marsh", Region: "East Basin",
		Latitude: 44.207, Longitude: 8.912, Habitats: []string{"wetland", "reedbed"},
	})
	_ = catalog.Species.Create(domain.Species{
		ID: "sp-robin", CommonName: "European robin", ScientificName: "Erithacus rubecula",
		Class: "bird", Tags: []string{"resident", "songbird"},
	})
	_ = catalog.Species.Create(domain.Species{
		ID: "sp-newt", CommonName: "Alpine newt", ScientificName: "Ichthyosaura alpestris",
		Class: "amphibian", Tags: []string{"wetland", "indicator"},
	})
	_ = catalog.Observations.Create(domain.Observation{
		ID: "obs-001", SiteID: "site-alpine", SpeciesID: "sp-robin",
		ObservedAt: time.Date(2026, 8, 17, 7, 15, 0, 0, time.UTC), Count: 3,
		Notes: "ridge trail observation", Status: domain.StatusApproved,
		CreatedAt: time.Date(2026, 8, 17, 7, 20, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 8, 17, 9, 0, 0, 0, time.UTC),
	})
	_ = catalog.Observations.Create(domain.Observation{
		ID: "obs-002", SiteID: "site-marsh", SpeciesID: "sp-newt",
		ObservedAt: time.Date(2026, 8, 17, 10, 45, 0, 0, time.UTC), Count: 5,
		Notes: "shallow pool", Status: domain.StatusDraft,
		CreatedAt: time.Date(2026, 8, 17, 10, 50, 0, 0, time.UTC),
		UpdatedAt: time.Date(2026, 8, 17, 10, 50, 0, 0, time.UTC),
	})
}
