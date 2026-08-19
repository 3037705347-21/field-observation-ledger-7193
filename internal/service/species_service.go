package service

import (
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
	"example.com/field-observation-ledger/internal/validate"
)

type SpeciesService struct {
	repository *repository.SpeciesRepository
}

func NewSpeciesService(repo *repository.SpeciesRepository) *SpeciesService {
	return &SpeciesService{repository: repo}
}

func (s *SpeciesService) Create(item domain.Species) (domain.Species, error) {
	if err := validate.Species(item); err != nil {
		return domain.Species{}, err
	}
	if err := s.repository.Create(item); err != nil {
		return domain.Species{}, err
	}
	// Return an isolated copy so the caller cannot mutate the stored record
	// through the returned Tags slice.
	return item.Clone(), nil
}

func (s *SpeciesService) Get(id string) (domain.Species, error) {
	item, err := s.repository.Get(id)
	if err != nil {
		return domain.Species{}, err
	}
	return item.Clone(), nil
}

func (s *SpeciesService) List() []domain.Species {
	items := s.repository.List()
	isolated := make([]domain.Species, len(items))
	for index := range items {
		isolated[index] = items[index].Clone()
	}
	return isolated
}

func (s *SpeciesService) Has(id string) bool {
	_, err := s.repository.Get(id)
	return err == nil
}
