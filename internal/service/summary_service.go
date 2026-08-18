package service

import (
	"example.com/field-observation-ledger/internal/analysis"
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

type SummaryService struct {
	observations *repository.ObservationRepository
}

func NewSummaryService(observations *repository.ObservationRepository) *SummaryService {
	return &SummaryService{observations: observations}
}

func (s *SummaryService) Build(filter repository.ObservationFilter) domain.Summary {
	summary := domain.NewSummary()
	items := s.observations.List(filter)
	for _, item := range items {
		summary.Add(item)
	}
	summary.Analysis = analysis.BuildReport(items, len(items))
	return summary
}

func (s *SummaryService) ForSite(siteID string) domain.Summary {
	return s.Build(repository.ObservationFilter{SiteID: siteID})
}

func (s *SummaryService) ForSpecies(speciesID string) domain.Summary {
	return s.Build(repository.ObservationFilter{SpeciesID: speciesID})
}
