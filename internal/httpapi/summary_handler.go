package httpapi

import (
	"net/http"

	"example.com/field-observation-ledger/internal/repository"
)

func (s *Server) summary(writer http.ResponseWriter, request *http.Request) {
	filter := repository.ObservationFilter{
		SiteID:    request.URL.Query().Get("site_id"),
		SpeciesID: request.URL.Query().Get("species_id"),
	}
	result := s.deps.Summary.Build(filter)
	writeSummary(writer, http.StatusOK, result)
}

func (s *Server) health(writer http.ResponseWriter, request *http.Request) {
	writeJSON(writer, http.StatusOK, map[string]string{"status": "ok", "service": "field-observation-ledger"})
}
