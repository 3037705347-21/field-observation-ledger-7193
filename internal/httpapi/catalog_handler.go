package httpapi

import (
	"errors"
	"net/http"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

func (s *Server) listSites(writer http.ResponseWriter, request *http.Request) {
	writeJSON(writer, http.StatusOK, map[string]any{"items": s.deps.Sites.List()})
}

func (s *Server) createSite(writer http.ResponseWriter, request *http.Request) {
	var item domain.Site
	if err := decodeJSON(request, &item); err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	created, err := s.deps.Sites.Create(item)
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, repository.ErrSiteExists) {
			status = http.StatusConflict
		}
		writeError(writer, status, err.Error())
		return
	}
	writeJSON(writer, http.StatusCreated, created)
}

func (s *Server) listSpecies(writer http.ResponseWriter, request *http.Request) {
	writeJSON(writer, http.StatusOK, map[string]any{"items": s.deps.Species.List()})
}

func (s *Server) createSpecies(writer http.ResponseWriter, request *http.Request) {
	var item domain.Species
	if err := decodeJSON(request, &item); err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	created, err := s.deps.Species.Create(item)
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, repository.ErrSpeciesExists) {
			status = http.StatusConflict
		}
		writeError(writer, status, err.Error())
		return
	}
	writeJSON(writer, http.StatusCreated, created)
}
