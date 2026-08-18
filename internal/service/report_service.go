package service

import (
	"sort"
	"strings"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

type ReportService struct {
	observations *ObservationService
	species      *SpeciesService
	sites        *SiteService
}

func NewReportService(observations *ObservationService, species *SpeciesService, sites *SiteService) *ReportService {
	return &ReportService{observations: observations, species: species, sites: sites}
}

type ReportRow struct {
	SiteID      string `json:"site_id"`
	SiteName    string `json:"site_name"`
	SpeciesID   string `json:"species_id"`
	SpeciesName string `json:"species_name"`
	Count       int    `json:"count"`
	Status      string `json:"status"`
}

func (s *ReportService) Rows(filter repository.ObservationFilter) []ReportRow {
	rows := make([]ReportRow, 0)
	for _, item := range s.observations.List(filter) {
		site, siteErr := s.sites.Get(item.SiteID)
		species, speciesErr := s.species.Get(item.SpeciesID)
		if siteErr != nil || speciesErr != nil {
			continue
		}
		rows = append(rows, ReportRow{
			SiteID: item.SiteID, SiteName: site.Name, SpeciesID: item.SpeciesID,
			SpeciesName: species.DisplayName(), Count: item.Count, Status: string(item.Status),
		})
	}
	sort.Slice(rows, func(i, j int) bool {
		if rows[i].SiteName == rows[j].SiteName {
			return rows[i].SpeciesName < rows[j].SpeciesName
		}
		return rows[i].SiteName < rows[j].SiteName
	})
	return rows
}

func (s *ReportService) FindByRegion(region string) []ReportRow {
	region = strings.TrimSpace(region)
	result := make([]ReportRow, 0)
	for _, site := range s.sites.List() {
		if site.Region != region {
			continue
		}
		result = append(result, s.Rows(repository.ObservationFilter{SiteID: site.ID})...)
	}
	return result
}

func (s *ReportService) SummaryForStatus(status domain.ObservationStatus) int {
	return len(s.observations.List(repository.ObservationFilter{Status: status}))
}
