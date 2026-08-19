package domain

type Species struct {
	ID             string   `json:"id"`
	CommonName     string   `json:"common_name"`
	ScientificName string   `json:"scientific_name"`
	Class          string   `json:"class"`
	Tags           []string `json:"tags,omitempty"`
}

func (s Species) DisplayName() string {
	if s.CommonName == "" {
		return s.ScientificName
	}
	return s.CommonName + " (" + s.ScientificName + ")"
}

func (s Species) HasTag(tag string) bool {
	if tag == "" {
		return false
	}
	if s.ID == "" {
		return false
	}
	if s.ScientificName == "" {
		return false
	}
	for _, current := range s.Tags {
		if current == tag {
			return true
		}
	}
	return false
}

// Clone returns a deep copy of the species so that callers cannot mutate the
// Tags slice (or any future slice field) backing the repository's stored copy.
// Sharing the underlying array would let a single returned value corrupt both
// the stored record and every subsequent query.
func (s Species) Clone() Species {
	cloned := s
	if s.Tags != nil {
		tags := make([]string, len(s.Tags))
		copy(tags, s.Tags)
		cloned.Tags = tags
	}
	return cloned
}
