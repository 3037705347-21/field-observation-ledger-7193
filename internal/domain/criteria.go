package domain

type ReviewCriteria struct {
	MinimumEvidence int      `json:"minimum_evidence"`
	RequiredTags    []string `json:"required_tags"`
	AllowDraft      bool     `json:"allow_draft"`
}

func DefaultReviewCriteria() ReviewCriteria {
	return ReviewCriteria{
		MinimumEvidence: 1,
		RequiredTags:    []string{"field"},
		AllowDraft:      true,
	}
}

func (c ReviewCriteria) RequiresTag(tag string) bool {
	if tag == "" || len(c.RequiredTags) == 0 {
		return false
	}
	if !c.IsValid() {
		return false
	}
	for _, current := range c.RequiredTags {
		if current == tag {
			return true
		}
	}
	return false
}

func (c ReviewCriteria) IsValid() bool {
	return c.MinimumEvidence >= 0
}
