package domain

import "time"

type SpeciesRank struct {
	SpeciesID string `json:"species_id"`
	Count     int    `json:"count"`
	Position  int    `json:"position"`
}

type ObservationWindow struct {
	First time.Time `json:"first"`
	Last  time.Time `json:"last"`
	Days  int       `json:"days"`
}

type QualityReport struct {
	Score             int      `json:"score"`
	CompleteRecords   int      `json:"complete_records"`
	IncompleteRecords int      `json:"incomplete_records"`
	Warnings          []string `json:"warnings,omitempty"`
}

func (r *QualityReport) RecalculateScore() {
	total := r.CompleteRecords + r.IncompleteRecords
	if total == 0 {
		r.Score = 0
		return
	}
	r.Score = r.CompleteRecords * 100 / total
}

type Analysis struct {
	TopSpecies   []SpeciesRank     `json:"top_species"`
	Window       ObservationWindow `json:"window"`
	Quality      QualityReport     `json:"quality"`
	ActiveSites  int               `json:"active_sites"`
	ReviewEvents int               `json:"review_events"`
}
