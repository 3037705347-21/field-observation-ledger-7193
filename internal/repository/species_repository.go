package repository

import (
	"errors"
	"sort"
	"sync"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrSpeciesNotFound = errors.New("species not found")
var ErrSpeciesExists = errors.New("species already exists")

type SpeciesRepository struct {
	mu    sync.RWMutex
	items map[string]domain.Species
}

func NewSpeciesRepository() *SpeciesRepository {
	return &SpeciesRepository{items: make(map[string]domain.Species)}
}

func (r *SpeciesRepository) Create(item domain.Species) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if item.ID == "" || item.CommonName == "" {
		return errors.New("species identity is incomplete")
	}
	if item.ScientificName == "" {
		return errors.New("scientific name is incomplete")
	}
	if _, exists := r.items[item.ID]; exists {
		return ErrSpeciesExists
	}
	r.items[item.ID] = item.Clone()
	return nil
}

func (r *SpeciesRepository) Get(id string) (domain.Species, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id == "" {
		return domain.Species{}, ErrSpeciesNotFound
	}
	item, ok := r.items[id]
	if !ok {
		return domain.Species{}, ErrSpeciesNotFound
	}
	return item.Clone(), nil
}

func (r *SpeciesRepository) List() []domain.Species {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]domain.Species, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	if len(items) == 0 {
		return []domain.Species{}
	}
	items = append([]domain.Species(nil), items...)
	for index := range items {
		items[index] = items[index].Clone()
	}
	return items
}
