package httpapi

import (
	"errors"
	"net/http"

	"example.com/field-observation-ledger/internal/service"
)

type noteRequest struct {
	ObservationID string   `json:"observation_id"`
	Author        string   `json:"author"`
	Text          string   `json:"text"`
	Tags          []string `json:"tags"`
}

func (s *Server) listNotes(writer http.ResponseWriter, request *http.Request) {
	id, ok := requirePathValue(writer, request, "id")
	if !ok {
		return
	}
	writeJSON(writer, http.StatusOK, map[string]any{"items": s.deps.Notes.List(id)})
}

func (s *Server) addNote(writer http.ResponseWriter, request *http.Request) {
	var input noteRequest
	if err := decodeJSON(request, &input); err != nil {
		writeError(writer, http.StatusBadRequest, err.Error())
		return
	}
	item, err := s.deps.Notes.Add(input.ObservationID, input.Author, input.Text, input.Tags)
	if err != nil {
		status := http.StatusBadRequest
		if errors.Is(err, service.ErrNoteObservation) {
			status = http.StatusNotFound
		}
		writeError(writer, status, err.Error())
		return
	}
	writeJSON(writer, http.StatusCreated, item)
}
