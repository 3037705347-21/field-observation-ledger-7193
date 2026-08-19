package analysis

import (
	"math"
	"strings"

	"example.com/field-observation-ledger/internal/domain"
)

func Quality(items []domain.Observation) domain.QualityReport {
	report := domain.QualityReport{}
	for _, item := range items {
		warnings := MissingFields(item)
		if len(warnings) == 0 {
			report.CompleteRecords++
			continue
		}
		report.IncompleteRecords++
		report.Warnings = appendUnique(report.Warnings, warnings...)
	}
	total := report.CompleteRecords + report.IncompleteRecords
	if total > 0 {
		report.Score = int(math.Round(float64(report.CompleteRecords) / float64(total) * 100))
	}
	return report
}

func MissingFields(item domain.Observation) []string {
	var warnings []string
	if strings.TrimSpace(item.ID) == "" {
		warnings = append(warnings, "missing observation id")
	}
	if strings.TrimSpace(item.SiteID) == "" {
		warnings = append(warnings, "missing site id")
	}
	if strings.TrimSpace(item.SpeciesID) == "" {
		warnings = append(warnings, "missing species id")
	}
	if item.ObservedAt.IsZero() {
		warnings = append(warnings, "missing observed time")
	}
	if item.Count <= 0 {
		warnings = append(warnings, "invalid count")
	}
	return warnings
}

func appendUnique(values []string, incoming ...string) []string {
	seen := make(map[string]bool, len(values)+len(incoming))
	for _, value := range values {
		seen[value] = true
	}
	for _, value := range incoming {
		if !seen[value] {
			seen[value] = true
			values = append(values, value)
		}
	}
	return values
}
