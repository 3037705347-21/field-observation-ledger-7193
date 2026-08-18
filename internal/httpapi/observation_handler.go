package httpapi

import (
	"errors"
	"net/http"
	"strings"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
	"example.com/field-observation-ledger/internal/service"
)

type observationRequest struct {
	ID         string `json:"id"`
	SiteID     string `json:"site_id"`
	SpeciesID  string `json:"species_id"`
	ObservedAt string `json:"observed_at"`
	Count      int    `json:"count"`
	Notes      string `json:"notes"`
}

func (s *Server) listObservations(writer http.ResponseWriter, request *http.Request) {
	filter := repository.ObservationFilter{
		SiteID:    request.URL.Query().Get("site_id"),
		SpeciesID: request.URL.Query().Get("species_id"),
		Status:    domain.ObservationStatus(request.URL.Query().Get("status")),
	}
	items := s.deps.Observations.List(filter)
	writeJSON(writer, http.StatusOK, map[string]any{
		"items": paginate(items, parsePage(request)),
		"total": len(items),
	})
}

func (s *Server) createObservation(writer http.ResponseWriter, request *http.Request) {
	var input observationRequest
	if err := decodeJSON(request, &input); err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	observedAt, err := service.ParseTime(input.ObservedAt)
	if err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	item, err := s.deps.Observations.Create(domain.Observation{
		ID: input.ID, SiteID: input.SiteID, SpeciesID: input.SpeciesID,
		ObservedAt: observedAt, Count: input.Count, Notes: strings.TrimSpace(input.Notes),
	})
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, repository.ErrObservationExists) {
			status = http.StatusConflict
		}
		writeError(writer, status, err.Error())
		return
	}
	writeJSON(writer, http.StatusCreated, item)
}
