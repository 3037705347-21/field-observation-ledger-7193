package service

import (
	"strings"

	"example.com/field-observation-ledger/internal/domain"
)

type CatalogService struct {
	species *SpeciesService
	sites   *SiteService
}

func NewCatalogService(species *SpeciesService, sites *SiteService) *CatalogService {
	return &CatalogService{species: species, sites: sites}
}

func (s *CatalogService) ValidateReferences(siteID, speciesID string) error {
	if _, err := s.sites.Get(strings.TrimSpace(siteID)); err != nil {
		return ErrSiteReference
	}
	if _, err := s.species.Get(strings.TrimSpace(speciesID)); err != nil {
		return ErrSpeciesReference
	}
	return nil
}

func (s *CatalogService) Lookup(siteID, speciesID string) (domain.Site, domain.Species, error) {
	site, err := s.sites.Get(siteID)
	if err != nil {
		return domain.Site{}, domain.Species{}, err
	}
	species, err := s.species.Get(speciesID)
	if err != nil {
		return domain.Site{}, domain.Species{}, err
	}
	return site, species, nil
}
