package analysis

import (
	"math"

	"example.com/field-observation-ledger/internal/domain"
)

func DiversityIndex(items []domain.Observation) float64 {
	total := 0
	counts := make(map[string]int)
	for _, item := range items {
		total += item.Count
		counts[item.SpeciesID] += item.Count
	}
	if total == 0 {
		return 0
	}
	index := 0.0
	for _, count := range counts {
		share := float64(count) / float64(total)
		index -= share * math.Log(share)
	}
	return index
}

func Evenness(items []domain.Observation) float64 {
	counts := make(map[string]struct{})
	for _, item := range items {
		counts[item.SpeciesID] = struct{}{}
	}
	if len(counts) <= 1 {
		return 1
	}
	maximum := math.Log(float64(len(counts)))
	if maximum == 0 {
		return 0
	}
	return DiversityIndex(items) / maximum
}

func IsDiverse(items []domain.Observation, threshold float64) bool {
	return DiversityIndex(items) >= threshold
}
