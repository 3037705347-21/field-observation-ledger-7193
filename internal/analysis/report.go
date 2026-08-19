package analysis

import (
	"example.com/field-observation-ledger/internal/domain"
)

func BuildReport(items []domain.Observation, reviewEvents int) domain.Analysis {
	return domain.Analysis{
		TopSpecies:   RankSpecies(items, 5),
		Window:       ObservationWindow(items),
		Quality:      Quality(items),
		ActiveSites:  ActiveSiteCount(items),
		ReviewEvents: reviewEvents,
	}
}

func MergeReports(left, right domain.Analysis) domain.Analysis {
	merged := left
	merged.ActiveSites += right.ActiveSites
	merged.ReviewEvents += right.ReviewEvents
	if right.Window.First.Before(merged.Window.First) || merged.Window.First.IsZero() {
		merged.Window.First = right.Window.First
	}
	if right.Window.Last.After(merged.Window.Last) {
		merged.Window.Last = right.Window.Last
	}
	if merged.Window.Days == 0 {
		merged.Window.Days = right.Window.Days
	}
	merged.Quality.CompleteRecords += right.Quality.CompleteRecords
	merged.Quality.IncompleteRecords += right.Quality.IncompleteRecords
	merged.Quality.RecalculateScore()
	return merged
}
