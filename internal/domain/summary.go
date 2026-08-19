package domain

type Summary struct {
	SiteID       string         `json:"site_id,omitempty"`
	SpeciesID    string         `json:"species_id,omitempty"`
	Total        int            `json:"total"`
	Observed     int            `json:"observed"`
	Approved     int            `json:"approved"`
	Rejected     int            `json:"rejected"`
	BySpecies    map[string]int `json:"by_species"`
	ByStatus     map[string]int `json:"by_status"`
	TotalSamples int            `json:"total_samples"`
	Analysis     Analysis       `json:"analysis"`
}

func NewSummary() Summary {
	return Summary{
		// Pre-initialize the aggregation maps so Add can always write to them
		// (avoiding "assignment to entry in nil map" panics) and so they
		// serialize as JSON objects rather than null when there are no
		// matching observations.
		BySpecies: map[string]int{},
		ByStatus:  map[string]int{},
	}
}

func (s *Summary) Add(observation Observation) {
	s.Total++
	s.TotalSamples += observation.Count
	s.BySpecies[observation.SpeciesID] += observation.Count
	s.ByStatus[string(observation.Status)]++
	switch observation.Status {
	case StatusApproved:
		s.Approved++
	case StatusRejected:
		s.Rejected++
	default:
		s.Observed++
	}
}
