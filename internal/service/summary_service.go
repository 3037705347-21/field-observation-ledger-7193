package service

import (
	"example.com/field-observation-ledger/internal/analysis"
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

type SummaryService struct {
	observations *repository.ObservationRepository
	reviews      *repository.ReviewRepository
}

func NewSummaryService(observations *repository.ObservationRepository, reviews ...*repository.ReviewRepository) *SummaryService {
	service := &SummaryService{observations: observations}
	if len(reviews) > 0 {
		service.reviews = reviews[0]
	}
	return service
}

func (s *SummaryService) Build(filter repository.ObservationFilter) domain.Summary {
	summary := domain.NewSummary()
	items := s.observations.List(filter)
	for _, item := range items {
		summary.Add(item)
	}
	reviewEvents := 0
	if s.reviews != nil {
		reviewEvents = s.reviews.Count()
	}
	summary.Analysis = analysis.BuildReport(items, reviewEvents)
	summary.Analysis.SetReviewEvents(reviewEvents)
	return summary
}

func (s *SummaryService) ForSite(siteID string) domain.Summary {
	return s.Build(repository.ObservationFilter{SiteID: siteID})
}

func (s *SummaryService) ForSpecies(speciesID string) domain.Summary {
	return s.Build(repository.ObservationFilter{SpeciesID: speciesID})
}
