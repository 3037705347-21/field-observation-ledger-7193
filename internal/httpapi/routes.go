package httpapi

import "net/http"

func (s *Server) registerRoutes() {
	s.mux.HandleFunc("GET /health", s.health)
	s.mux.HandleFunc("GET /api/v1/sites", s.listSites)
	s.mux.HandleFunc("POST /api/v1/sites", s.createSite)
	s.mux.HandleFunc("GET /api/v1/species", s.listSpecies)
	s.mux.HandleFunc("POST /api/v1/species", s.createSpecies)
	s.mux.HandleFunc("GET /api/v1/observations", s.listObservations)
	s.mux.HandleFunc("POST /api/v1/observations", s.createObservation)
	s.mux.HandleFunc("POST /api/v1/reviews/{id}/approve", s.approveObservation)
	s.mux.HandleFunc("POST /api/v1/reviews/{id}/reject", s.rejectObservation)
	s.mux.HandleFunc("GET /api/v1/reviews/{id}", s.listReviews)
	s.mux.HandleFunc("GET /api/v1/evidence/{id}", s.listEvidence)
	s.mux.HandleFunc("POST /api/v1/evidence", s.addEvidence)
	s.mux.HandleFunc("POST /api/v1/evidence/{id}/verify", s.verifyEvidence)
	s.mux.HandleFunc("GET /api/v1/notes/{id}", s.listNotes)
	s.mux.HandleFunc("POST /api/v1/notes", s.addNote)
	s.mux.HandleFunc("GET /api/v1/summary", s.summary)
	s.mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writeError(writer, http.StatusNotFound, "route not found")
	})
}
