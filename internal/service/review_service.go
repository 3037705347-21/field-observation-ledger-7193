package service

import (
	"errors"
	"fmt"
	"strings"

	"example.com/field-observation-ledger/internal/clock"
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

var ErrReviewerRequired = errors.New("reviewer is required")
var ErrDecisionRequired = errors.New("decision is required")
var ErrDecisionInvalid = errors.New("decision must be approve or reject")

type ReviewService struct {
	observations *ObservationService
	reviews      *repository.ReviewRepository
	clock        clock.Clock
}

func NewReviewService(observations *ObservationService, reviews *repository.ReviewRepository, timeSource clock.Clock) *ReviewService {
	return &ReviewService{observations: observations, reviews: reviews, clock: timeSource}
}

func (s *ReviewService) Decide(observationID, reviewer string, decision domain.ReviewDecision, comment string) (domain.Observation, domain.Review, error) {
	if strings.TrimSpace(reviewer) == "" {
		return domain.Observation{}, domain.Review{}, ErrReviewerRequired
	}
	if decision == "" {
		return domain.Observation{}, domain.Review{}, ErrDecisionRequired
	}
	if decision != domain.DecisionApprove && decision != domain.DecisionReject {
		return domain.Observation{}, domain.Review{}, ErrDecisionInvalid
	}
	status := domain.StatusApproved
	if decision == domain.DecisionReject {
		status = domain.StatusRejected
	}
	observation, err := s.observations.ChangeStatus(observationID, status)
	if err != nil {
		return domain.Observation{}, domain.Review{}, err
	}
	review := domain.Review{
		ID:            fmt.Sprintf("review-%s-%d", observationID, s.clock.Now().UnixNano()),
		ObservationID: observationID,
		Reviewer:      reviewer,
		Decision:      decision,
		Comment:       strings.TrimSpace(comment),
		CreatedAt:     s.clock.Now(),
	}
	s.reviews.Create(review)
	return observation, review, nil
}

func (s *ReviewService) History(observationID string) []domain.Review {
	return s.reviews.ListForObservation(observationID)
}
