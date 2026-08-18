package service

import (
	"errors"
	"strings"

	"example.com/field-observation-ledger/internal/clock"
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/metrics"
	"example.com/field-observation-ledger/internal/repository"
	"example.com/field-observation-ledger/internal/validate"
)

var ErrSiteReference = errors.New("site reference does not exist")
var ErrSpeciesReference = errors.New("species reference does not exist")
var ErrObservationState = errors.New("observation is not reviewable")

type ObservationService struct {
	observations *repository.ObservationRepository
	sites        *repository.SiteRepository
	species      *repository.SpeciesRepository
	clock        clock.Clock
	metrics      *metrics.Counter
}

func NewObservationService(
	observations *repository.ObservationRepository,
	sites *repository.SiteRepository,
	species *repository.SpeciesRepository,
	timeSource clock.Clock,
	counter *metrics.Counter,
) *ObservationService {
	return &ObservationService{
		observations: observations,
		sites:        sites,
		species:      species,
		clock:        timeSource,
		metrics:      counter,
	}
}

func (s *ObservationService) Create(item domain.Observation) (domain.Observation, error) {
	if err := validate.Observation(item); err != nil {
		return domain.Observation{}, err
	}
	if _, err := s.sites.Get(item.SiteID); err != nil {
		return domain.Observation{}, ErrSiteReference
	}
	if _, err := s.species.Get(item.SpeciesID); err != nil {
		return domain.Observation{}, ErrSpeciesReference
	}
	now := s.clock.Now()
	item.Status = domain.StatusDraft
	item.CreatedAt = now
	item.UpdatedAt = now
	if err := s.observations.Create(item); err != nil {
		return domain.Observation{}, err
	}
	s.metrics.Add("observations.created", 1)
	return item, nil
}

func (s *ObservationService) Get(id string) (domain.Observation, error) {
	return s.observations.Get(strings.TrimSpace(id))
}

func (s *ObservationService) List(filter repository.ObservationFilter) []domain.Observation {
	return s.observations.List(filter)
}

func (s *ObservationService) ChangeStatus(id string, status domain.ObservationStatus) (domain.Observation, error) {
	item, err := s.Get(id)
	if err != nil {
		return domain.Observation{}, err
	}
	if !item.IsReviewable() {
		return domain.Observation{}, ErrObservationState
	}
	item.Status = status
	item.UpdatedAt = s.clock.Now()
	if err := s.observations.Update(item); err != nil {
		return domain.Observation{}, err
	}
	s.metrics.Add("observations.reviewed", 1)
	return item, nil
}
