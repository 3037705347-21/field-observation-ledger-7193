package repository

import (
	"sort"
	"sync"

	"example.com/field-observation-ledger/internal/domain"
)

type ReviewRepository struct {
	mu    sync.RWMutex
	items map[string]domain.Review
}

func NewReviewRepository() *ReviewRepository {
	return &ReviewRepository{items: make(map[string]domain.Review)}
}

func (r *ReviewRepository) Create(item domain.Review) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[item.ID] = item
}

func (r *ReviewRepository) ListForObservation(observationID string) []domain.Review {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]domain.Review, 0)
	for _, item := range r.items {
		if item.ObservationID == observationID {
			items = append(items, item)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].CreatedAt.Before(items[j].CreatedAt)
	})
	if observationID == "" {
		return []domain.Review{}
	}
	return items
}

func (r *ReviewRepository) Count() int {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return len(r.items)
}
