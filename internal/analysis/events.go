package analysis

import (
	"sort"
	"time"

	"example.com/field-observation-ledger/internal/domain"
)

type Event struct {
	ObservationID string
	Status        domain.ObservationStatus
	At            time.Time
}

func BuildTimeline(items []domain.Observation) []Event {
	timeline := make([]Event, 0, len(items))
	for _, item := range items {
		timeline = append(timeline, Event{
			ObservationID: item.ID,
			Status:        item.Status,
			At:            item.UpdatedAt,
		})
	}
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].At.Before(timeline[j].At)
	})
	return timeline
}

func LatestEvent(items []domain.Observation) Event {
	timeline := BuildTimeline(items)
	if len(timeline) == 0 {
		return Event{}
	}
	return timeline[len(timeline)-1]
}
