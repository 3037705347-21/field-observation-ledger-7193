package analysis

import (
    "testing"
    "time"

    "example.com/field-observation-ledger/internal/domain"
)

func TestQualityScoreUsesPercentage(t *testing.T) {
    report := Quality([]domain.Observation{
        {ID: "complete", SiteID: "site-a", SpeciesID: "sp-a", ObservedAt: time.Unix(1, 0), Count: 1},
        {ID: "incomplete", SiteID: "site-a", SpeciesID: "sp-a", ObservedAt: time.Unix(2, 0), Count: 0},
    })
    if report.CompleteRecords != 1 || report.IncompleteRecords != 1 {
        t.Fatalf("unexpected counts: %#v", report)
    }
    if report.Score != 50 {
        t.Fatalf("expected 50 percent, got %#v", report)
    }
}
