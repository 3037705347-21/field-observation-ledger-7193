package httpapi

import (
	"errors"
	"net/http"
	"time"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
	"example.com/field-observation-ledger/internal/service"
)

type evidenceRequest struct {
	ID            string `json:"id"`
	ObservationID string `json:"observation_id"`
	Kind          string `json:"kind"`
	Reference     string `json:"reference"`
	CapturedAt    string `json:"captured_at"`
	Observer      string `json:"observer"`
}

func (s *Server) listEvidence(writer http.ResponseWriter, request *http.Request) {
	id, ok := requirePathValue(writer, request, "id")
	if !ok {
		return
	}
	writeJSON(writer, http.StatusOK, map[string]any{"items": s.deps.Evidence.List(id)})
}

func (s *Server) addEvidence(writer http.ResponseWriter, request *http.Request) {
	var input evidenceRequest
	if err := decodeJSON(request, &input); err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	capturedAt, err := parseOptionalTime(input.CapturedAt)
	if err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	item, err := s.deps.Evidence.Add(domain.Evidence{
		ID: input.ID, ObservationID: input.ObservationID,
		Kind: domain.EvidenceKind(input.Kind), Reference: input.Reference,
		CapturedAt: capturedAt, Observer: input.Observer,
	})
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, repository.ErrEvidenceExists) {
			status = http.StatusConflict
		}
		if errors.Is(err, service.ErrEvidenceObservation) {
			status = http.StatusNotFound
		}
		writeError(writer, status, err.Error())
		return
	}
	writeJSON(writer, http.StatusCreated, item)
}

func (s *Server) verifyEvidence(writer http.ResponseWriter, request *http.Request) {
	id, ok := requirePathValue(writer, request, "id")
	if !ok {
		return
	}
	item, err := s.deps.Evidence.Verify(id)
	if err != nil {
		writeError(writer, http.StatusNotFound, err.Error())
		return
	}
	writeJSON(writer, http.StatusOK, item)
}

func parseOptionalTime(value string) (time.Time, error) {
	if value == "" {
		return time.Time{}, nil
	}
	return service.ParseTime(value)
}
