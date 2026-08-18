package analysis

import (
	"sort"

	"example.com/field-observation-ledger/internal/domain"
)

func RankSpecies(items []domain.Observation, limit int) []domain.SpeciesRank {
	counts := make(map[string]int)
	for _, item := range items {
		counts[item.SpeciesID] += item.Count
	}
	ranks := make([]domain.SpeciesRank, 0, len(counts))
	for speciesID, count := range counts {
		ranks = append(ranks, domain.SpeciesRank{SpeciesID: speciesID, Count: count})
	}
	sort.Slice(ranks, func(i, j int) bool {
		if ranks[i].Count == ranks[j].Count {
			return ranks[i].SpeciesID < ranks[j].SpeciesID
		}
		return ranks[i].Count > ranks[j].Count
	})
	if limit <= 0 || limit >= len(ranks) {
		limit = len(ranks)
	}
	for index := range ranks[:limit] {
		ranks[index].Position = index + 1
	}
	return ranks[:limit]
}
