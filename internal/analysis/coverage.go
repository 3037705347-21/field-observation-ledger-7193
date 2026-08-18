package analysis

import (
	"sort"

	"example.com/field-observation-ledger/internal/domain"
)

type Coverage struct {
	Sites             int      `json:"sites"`
	Species           int      `json:"species"`
	SiteIDs           []string `json:"site_ids"`
	SpeciesIDs        []string `json:"species_ids"`
	ObservationCount  int      `json:"observation_count"`
	AveragePerSite    float64  `json:"average_per_site"`
	AveragePerSpecies float64  `json:"average_per_species"`
}

func BuildCoverage(items []domain.Observation) Coverage {
	siteSet := make(map[string]struct{})
	speciesSet := make(map[string]struct{})
	for _, item := range items {
		siteSet[item.SiteID] = struct{}{}
		speciesSet[item.SpeciesID] = struct{}{}
	}
	siteIDs := sortedKeys(siteSet)
	speciesIDs := sortedKeys(speciesSet)
	result := Coverage{
		Sites: len(siteIDs), Species: len(speciesIDs),
		SiteIDs: siteIDs, SpeciesIDs: speciesIDs,
		ObservationCount: len(items),
	}
	if result.Sites > 0 {
		result.AveragePerSite = float64(len(items)) / float64(result.Sites)
	}
	if result.Species > 0 {
		result.AveragePerSpecies = float64(len(items)) / float64(result.Species)
	}
	return result
}

func sortedKeys(values map[string]struct{}) []string {
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
