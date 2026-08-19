package httpapi

import (
	"net/http"
	"example.com/field-observation-ledger/internal/config"
	"example.com/field-observation-ledger/internal/service"
)

type Dependencies struct {
	Observations *service.ObservationService
	Species *service.SpeciesService
	Sites *service.SiteService
	Reviews *service.ReviewService
	Summary *service.SummaryService
	Evidence *service.EvidenceService
	Notes *service.NoteService
	PageLimit int
}
type Server struct { deps Dependencies; mux *http.ServeMux }
func NewServer(deps Dependencies) *Server {
	if deps.PageLimit <= 0 { deps.PageLimit = config.DefaultPageSize }
	server := &Server{deps: deps, mux: http.NewServeMux()}
	server.registerRoutes()
	return server
}
func (s *Server) Handler() http.Handler { return loggingMiddleware(jsonMiddleware(s.mux)) }
