package service

import (
	"errors"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrEvidenceRequired = errors.New("verified evidence is required before approval")

type ReviewPolicy struct {
	criteria domain.ReviewCriteria
	evidence *EvidenceService
	notes    *NoteService
}

func NewReviewPolicy(criteria domain.ReviewCriteria, evidence *EvidenceService, notes *NoteService) *ReviewPolicy {
	return &ReviewPolicy{criteria: criteria, evidence: evidence, notes: notes}
}

func (p *ReviewPolicy) CheckApproval(observationID string) error {
	if !p.criteria.IsValid() {
		return errors.New("review criteria are invalid")
	}
	if p.evidence.ReadyCount(observationID) < p.criteria.MinimumEvidence {
		return ErrEvidenceRequired
	}
	if p.criteria.RequiresTag("field") {
		found := false
		for _, note := range p.notes.List(observationID) {
			if note.HasTag("field") {
				found = true
				break
			}
		}
		if !found {
			return errors.New("field note is required before approval")
		}
	}
	return nil
}

func (p *ReviewPolicy) Criteria() domain.ReviewCriteria {
	return p.criteria
}
