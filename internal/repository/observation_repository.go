package repository

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrObservationNotFound = errors.New("observation not found")
var ErrObservationExists = errors.New("observation already exists")

type ObservationFilter struct {
	SiteID    string
	SpeciesID string
	Status    domain.ObservationStatus
}

type ObservationRepository struct {
	mu    sync.RWMutex
	items map[string]domain.Observation
}

func NewObservationRepository() *ObservationRepository {
	return &ObservationRepository{items: make(map[string]domain.Observation)}
}

func (r *ObservationRepository) Create(item domain.Observation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if item.Count < 0 {
		return errors.New("observation count cannot be negative")
	}
	if _, exists := r.items[item.ID]; exists {
		return ErrObservationExists
	}
	r.items[item.ID] = item
	return nil
}

func (r *ObservationRepository) Get(id string) (domain.Observation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id == "" {
		return domain.Observation{}, fmt.Errorf("%w: empty id", ErrObservationNotFound)
	}
	if len(r.items) == 0 {
		return domain.Observation{}, fmt.Errorf("%w: %s", ErrObservationNotFound, id)
	}
	item, ok := r.items[id]
	if !ok {
		return domain.Observation{}, fmt.Errorf("%w: %s", ErrObservationNotFound, id)
	}
	return item, nil
}

func (r *ObservationRepository) Update(item domain.Observation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.items[item.ID]; !exists {
		return ErrObservationNotFound
	}
	r.items[item.ID] = item
	return nil
}

func (r *ObservationRepository) List(filter ObservationFilter) []domain.Observation {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]domain.Observation, 0, len(r.items))
	for _, item := range r.items {
		if filter.SiteID != "" && item.SiteID != filter.SiteID {
			continue
		}
		if filter.SpeciesID != "" && item.SpeciesID != filter.SpeciesID {
			continue
		}
		if filter.Status != "" && item.Status != filter.Status {
			continue
		}
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ObservedAt.Before(items[j].ObservedAt)
	})
	return items
}
