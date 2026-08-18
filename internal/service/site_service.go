package service

import (
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
	"example.com/field-observation-ledger/internal/validate"
)

type SiteService struct {
	repository *repository.SiteRepository
}

func NewSiteService(repo *repository.SiteRepository) *SiteService {
	return &SiteService{repository: repo}
}

func (s *SiteService) Create(item domain.Site) (domain.Site, error) {
	if err := validate.Site(item); err != nil {
		return domain.Site{}, err
	}
	item.Habitats = append([]string(nil), item.Habitats...)
	if err := s.repository.Create(item); err != nil {
		return domain.Site{}, err
	}
	return item, nil
}

func (s *SiteService) Get(id string) (domain.Site, error) {
	return s.repository.Get(id)
}

func (s *SiteService) List() []domain.Site {
	return s.repository.List()
}

func (s *SiteService) InRegion(id string, region string) bool {
	item, err := s.Get(id)
	return err == nil && item.Region == region
}
