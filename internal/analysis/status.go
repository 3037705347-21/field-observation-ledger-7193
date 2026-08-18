package analysis

import "example.com/field-observation-ledger/internal/domain"

func StatusCounts(items []domain.Observation) map[domain.ObservationStatus]int {
	counts := map[domain.ObservationStatus]int{
		domain.StatusDraft:    0,
		domain.StatusApproved: 0,
		domain.StatusRejected: 0,
	}
	for _, item := range items {
		counts[item.Status]++
	}
	return counts
}

func ApprovedShare(items []domain.Observation) float64 {
	if len(items) == 0 {
		return 0
	}
	return float64(StatusCounts(items)[domain.StatusApproved]) / float64(len(items))
}

func HasOpenReview(items []domain.Observation) bool {
	for _, item := range items {
		if item.Status == domain.StatusDraft {
			return true
		}
	}
	return false
}
