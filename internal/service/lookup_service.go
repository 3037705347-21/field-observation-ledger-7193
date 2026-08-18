package service

import (
	"strings"

	"example.com/field-observation-ledger/internal/domain"
)

type LookupService struct {
	species *SpeciesService
	sites   *SiteService
}

func NewLookupService(species *SpeciesService, sites *SiteService) *LookupService {
	return &LookupService{species: species, sites: sites}
}

type LookupResult struct {
	Site      domain.Site    `json:"site"`
	Species   domain.Species `json:"species"`
	FoundSite bool           `json:"found_site"`
	FoundSpec bool           `json:"found_species"`
}

func (s *LookupService) Resolve(siteID, speciesID string) LookupResult {
	result := LookupResult{}
	if siteID != "" {
		if site, err := s.sites.Get(strings.TrimSpace(siteID)); err == nil {
			result.Site = site
			result.FoundSite = true
		}
	}
	if speciesID != "" {
		if species, err := s.species.Get(strings.TrimSpace(speciesID)); err == nil {
			result.Species = species
			result.FoundSpec = true
		}
	}
	return result
}

func (s *LookupService) SearchSpecies(term string) []domain.Species {
	term = strings.ToLower(strings.TrimSpace(term))
	result := make([]domain.Species, 0)
	for _, item := range s.species.List() {
		if term == "" || strings.Contains(strings.ToLower(item.DisplayName()), term) {
			result = append(result, item)
		}
	}
	return result
}

func (s *LookupService) SearchSites(term string) []domain.Site {
	term = strings.ToLower(strings.TrimSpace(term))
	result := make([]domain.Site, 0)
	for _, item := range s.sites.List() {
		if term == "" || strings.Contains(strings.ToLower(item.Name), term) ||
			strings.Contains(strings.ToLower(item.Region), term) {
			result = append(result, item)
		}
	}
	return result
}
