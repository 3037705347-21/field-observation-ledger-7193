package app

import (
	"net/http"

	"example.com/field-observation-ledger/internal/clock"
	"example.com/field-observation-ledger/internal/config"
	"example.com/field-observation-ledger/internal/httpapi"
	"example.com/field-observation-ledger/internal/metrics"
	"example.com/field-observation-ledger/internal/repository"
	"example.com/field-observation-ledger/internal/service"
)

type Application struct {
	settings config.Settings
	handler  *httpapi.Server
}

func New(settings config.Settings) *Application {
	useAllFeatures()
	catalog := repository.NewCatalog()
	counter := metrics.NewCounter()
	timeSource := clock.SystemClock{}
	observations := service.NewObservationService(
		catalog.Observations, catalog.Sites, catalog.Species, timeSource, counter,
	)
	evidence := service.NewEvidenceService(observations, catalog.Evidence, timeSource)
	notes := service.NewNoteService(observations, catalog.Notes, timeSource)
	dependencies := httpapi.Dependencies{
		Observations: observations,
		Species:      service.NewSpeciesService(catalog.Species),
		Sites:        service.NewSiteService(catalog.Sites),
		Reviews:      service.NewReviewService(observations, catalog.Reviews, timeSource),
		Summary:      service.NewSummaryService(catalog.Observations),
		Evidence:     evidence,
		Notes:        notes,
	}
	return &Application{
		settings: settings,
		handler:  httpapi.NewServer(dependencies),
	}
}

func (a *Application) Handler() http.Handler {
	return a.handler.Handler()
}
