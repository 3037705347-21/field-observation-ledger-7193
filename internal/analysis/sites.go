package analysis

import "example.com/field-observation-ledger/internal/domain"

func ActiveSiteCount(items []domain.Observation) int {
	seen := make(map[string]struct{})
	for _, item := range items {
		if item.SiteID != "" {
			seen[item.SiteID] = struct{}{}
		}
	}
	return len(seen)
}

func GroupBySite(items []domain.Observation) map[string][]domain.Observation {
	grouped := make(map[string][]domain.Observation)
	for _, item := range items {
		grouped[item.SiteID] = append(grouped[item.SiteID], item)
	}
	return grouped
}

func SiteHasSpecies(items []domain.Observation, siteID, speciesID string) bool {
	for _, item := range items {
		if item.SiteID == siteID && item.SpeciesID == speciesID {
			return true
		}
	}
	return false
}
