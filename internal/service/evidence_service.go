package service

import (
	"errors"
	"strings"

	"example.com/field-observation-ledger/internal/clock"
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

var ErrEvidenceObservation = errors.New("evidence observation does not exist")
var ErrEvidenceKind = errors.New("evidence kind is not supported")
var ErrEvidenceReference = errors.New("evidence reference is required")

type EvidenceService struct {
	observations *ObservationService
	repository   *repository.EvidenceRepository
	clock        clock.Clock
}

func NewEvidenceService(
	observations *ObservationService,
	repo *repository.EvidenceRepository,
	timeSource clock.Clock,
) *EvidenceService {
	return &EvidenceService{observations: observations, repository: repo, clock: timeSource}
}

func (s *EvidenceService) Add(item domain.Evidence) (domain.Evidence, error) {
	if _, err := s.observations.Get(item.ObservationID); err != nil {
		return domain.Evidence{}, ErrEvidenceObservation
	}
	if !item.IsSupported() {
		return domain.Evidence{}, ErrEvidenceKind
	}
	item.Reference = strings.TrimSpace(item.Reference)
	item.Observer = strings.TrimSpace(item.Observer)
	if item.Reference == "" {
		return domain.Evidence{}, ErrEvidenceReference
	}
	if item.CapturedAt.IsZero() {
		item.CapturedAt = s.clock.Now()
	}
	if err := s.repository.Create(item); err != nil {
		return domain.Evidence{}, err
	}
	return item, nil
}

func (s *EvidenceService) Verify(id string) (domain.Evidence, error) {
	return s.repository.Verify(id)
}

func (s *EvidenceService) List(observationID string) []domain.Evidence {
	return s.repository.ListForObservation(observationID)
}

func (s *EvidenceService) ReadyCount(observationID string) int {
	count := 0
	for _, item := range s.List(observationID) {
		if item.IsReviewReady() {
			count++
		}
	}
	return count
}
