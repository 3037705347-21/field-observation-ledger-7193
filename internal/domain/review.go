package domain

import "time"

type ReviewDecision string

const (
	DecisionApprove ReviewDecision = "approve"
	DecisionReject  ReviewDecision = "reject"
)

type Review struct {
	ID            string         `json:"id"`
	ObservationID string         `json:"observation_id"`
	Reviewer      string         `json:"reviewer"`
	Decision      ReviewDecision `json:"decision"`
	Comment       string         `json:"comment"`
	CreatedAt     time.Time      `json:"created_at"`
}

func (r Review) IsApproval() bool {
	return r.Decision == DecisionApprove
}

func (r Review) IsRejection() bool {
	return r.Decision == DecisionReject
}
