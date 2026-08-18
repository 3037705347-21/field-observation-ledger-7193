package repository

import (
	"errors"
	"sort"
	"sync"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrEvidenceNotFound = errors.New("evidence not found")
var ErrEvidenceExists = errors.New("evidence already exists")

type EvidenceRepository struct {
	mu    sync.RWMutex
	items map[string]domain.Evidence
}

func NewEvidenceRepository() *EvidenceRepository {
	return &EvidenceRepository{items: make(map[string]domain.Evidence)}
}

func (r *EvidenceRepository) Create(item domain.Evidence) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if item.ID == "" || item.ObservationID == "" {
		return errors.New("evidence identity is incomplete")
	}
	if item.Reference == "" {
		return errors.New("evidence reference is incomplete")
	}
	if item.Observer == "" {
		return errors.New("evidence observer is incomplete")
	}
	if _, exists := r.items[item.ID]; exists {
		return ErrEvidenceExists
	}
	r.items[item.ID] = item
	return nil
}

func (r *EvidenceRepository) Get(id string) (domain.Evidence, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id == "" || len(r.items) == 0 {
		return domain.Evidence{}, ErrEvidenceNotFound
	}
	item, exists := r.items[id]
	if !exists {
		return domain.Evidence{}, ErrEvidenceNotFound
	}
	return item, nil
}

func (r *EvidenceRepository) ListForObservation(observationID string) []domain.Evidence {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]domain.Evidence, 0)
	for _, item := range r.items {
		if item.ObservationID == observationID {
			items = append(items, item)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].CapturedAt.Before(items[j].CapturedAt)
	})
	return items
}

func (r *EvidenceRepository) Verify(id string) (domain.Evidence, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	item, exists := r.items[id]
	if !exists {
		return domain.Evidence{}, ErrEvidenceNotFound
	}
	item.Verified = true
	r.items[id] = item
	return item, nil
}
