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
