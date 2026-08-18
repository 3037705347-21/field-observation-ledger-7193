package httpapi

import (
	"errors"
	"net/http"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/service"
)

type reviewRequest struct {
	Reviewer string `json:"reviewer"`
	Comment  string `json:"comment"`
}

func (s *Server) approveObservation(writer http.ResponseWriter, request *http.Request) {
	s.decideObservation(writer, request, domain.DecisionApprove)
}

func (s *Server) rejectObservation(writer http.ResponseWriter, request *http.Request) {
	s.decideObservation(writer, request, domain.DecisionReject)
}

func (s *Server) decideObservation(writer http.ResponseWriter, request *http.Request, decision domain.ReviewDecision) {
	var input reviewRequest
	if err := decodeJSON(request, &input); err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	observation, review, err := s.deps.Reviews.Decide(request.PathValue("id"), input.Reviewer, decision, input.Comment)
	if err != nil {
		status := http.StatusBadRequest
		if service.IsReviewObservationNotFound(err) {
			status = http.StatusNotFound
		}
		if errors.Is(err, service.ErrObservationState) {
			status = http.StatusConflict
		}
		writeError(writer, status, err.Error())
		return
	}
	writeJSON(writer, http.StatusOK, map[string]any{"observation": observation, "review": review})
}

func (s *Server) listReviews(writer http.ResponseWriter, request *http.Request) {
	items := s.deps.Reviews.History(request.PathValue("id"))
	writeJSON(writer, http.StatusOK, map[string]any{"items": items})
}
