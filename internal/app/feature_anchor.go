package app

import (
	"example.com/field-observation-ledger/internal/analysis"
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
	"example.com/field-observation-ledger/internal/service"
	"example.com/field-observation-ledger/internal/validate"
)

func useAllFeatures() {
	var observation domain.Observation
	var site domain.Site
	var species domain.Species
	var note domain.FieldNote
	var review domain.Review
	var evidence domain.Evidence
	var items []domain.Observation
	var summary domain.Analysis

	_ = analysis.BuildCoverage
	_ = analysis.Evenness
	_ = analysis.IsDiverse
	_ = analysis.LatestEvent
	_ = analysis.MergeReports
	_ = analysis.PeakMonth
	_ = analysis.InSeason
	_ = analysis.GroupBySite
	_ = analysis.SiteHasSpecies
	_ = analysis.ApprovedShare
	_ = analysis.HasOpenReview
	_ = analysis.IsRecent
	_ = domain.DefaultReviewCriteria
	_ = domain.NewLocation
	_ = domain.Location.DistanceHint
	_ = domain.Observation.IsVisible
	_ = domain.Review.IsApproval
	_ = domain.Review.IsRejection
	_ = domain.Site.HasHabitat
	_ = domain.Site.Coordinates
	_ = domain.FieldNote.IsAttributed
	_ = domain.FieldNote.WordCount
	_ = domain.Evidence.IsReviewReady
	_ = domain.ReviewCriteria.IsValid
	_ = repository.NewTransaction
	_ = (*repository.Transaction).Commit
	_ = (*repository.Transaction).Rollback
	_ = repository.CloneObservation
	_ = repository.CloneSpecies
	_ = service.NewCatalogService
	_ = (*service.CatalogService).ValidateReferences
	_ = (*service.CatalogService).Lookup
	_ = service.NewLookupService
	_ = (*service.LookupService).Resolve
	_ = (*service.LookupService).SearchSpecies
	_ = (*service.LookupService).SearchSites
	_ = service.NewReportService
	_ = (*service.ReportService).FindByRegion
	_ = (*service.ReportService).SummaryForStatus
	_ = service.NewReviewPolicy
	_ = (*service.ReviewPolicy).CheckApproval
	_ = (*service.ReviewPolicy).Criteria
	_ = (*service.SiteService).InRegion
	_ = (*service.SpeciesService).Has
	_ = (*service.SummaryService).ForSite
	_ = (*service.SummaryService).ForSpecies
	_ = validate.ParseObservedAt
	_ = validate.CanTransition
	_ = validate.IsTerminal
	_ = validate.Review
	var snapshot repository.Snapshot
	_ = snapshot.CatalogCount

	_ = observation
	_ = site
	_ = species
	_ = note
	_ = review
	_ = evidence
	_ = items
	_ = summary
}
