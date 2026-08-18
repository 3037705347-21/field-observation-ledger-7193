package validate

import (
	"errors"
	"strings"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrReviewID = errors.New("observation id is required")
var ErrReviewComment = errors.New("review comment is required")
var ErrReviewDecision = errors.New("review decision is invalid")

func Review(observationID, reviewer, comment string, decision domain.ReviewDecision) error {
	if strings.TrimSpace(observationID) == "" || strings.TrimSpace(reviewer) == "" {
		return ErrReviewID
	}
	if strings.TrimSpace(comment) == "" {
		return ErrReviewComment
	}
	if decision != domain.DecisionApprove && decision != domain.DecisionReject {
		return ErrReviewDecision
	}
	return nil
}

func CanTransition(from, to domain.ObservationStatus) bool {
	if from != domain.StatusDraft {
		return false
	}
	return to == domain.StatusApproved || to == domain.StatusRejected
}

func IsTerminal(status domain.ObservationStatus) bool {
	return status == domain.StatusApproved || status == domain.StatusRejected
}
