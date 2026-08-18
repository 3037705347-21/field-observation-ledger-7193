package analysis

import (
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

func MonthCounts(items []domain.Observation) map[time.Month]int {
	counts := make(map[time.Month]int)
	for _, item := range items {
		counts[item.ObservedAt.Month()] += item.Count
	}
	return counts
}

func PeakMonth(items []domain.Observation) time.Month {
	counts := MonthCounts(items)
	var peak time.Month
	best := 0
	for month, count := range counts {
		if count > best || (count == best && month < peak) {
			best = count
			peak = month
		}
	}
	return peak
}

func InSeason(item domain.Observation, start, end time.Month) bool {
	month := item.ObservedAt.Month()
	if start <= end {
		return month >= start && month <= end
	}
	return month >= start || month <= end
}
