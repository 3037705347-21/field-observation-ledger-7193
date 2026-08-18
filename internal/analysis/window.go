package analysis

import (
	"sort"
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

func ObservationWindow(items []domain.Observation) domain.ObservationWindow {
	if len(items) == 0 {
		return domain.ObservationWindow{}
	}
	dates := make([]time.Time, 0, len(items))
	for _, item := range items {
		dates = append(dates, item.ObservedAt)
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	first := dates[0]
	last := dates[len(dates)-1]
	days := int(last.Sub(first).Hours()/24) + 1
	if days < 1 {
		days = 1
	}
	return domain.ObservationWindow{First: first, Last: last, Days: days}
}

func IsRecent(item domain.Observation, now time.Time, days int) bool {
	if days <= 0 {
		return false
	}
	cutoff := now.Add(-time.Duration(days) * 24 * time.Hour)
	return !item.ObservedAt.Before(cutoff) && !item.ObservedAt.After(now)
}
