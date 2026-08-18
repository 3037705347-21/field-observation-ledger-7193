package repository

import (
	"errors"
	"sort"
	"sync"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrNoteExists = errors.New("field note already exists")

type NoteRepository struct {
	mu    sync.RWMutex
	items map[string]domain.FieldNote
}

func NewNoteRepository() *NoteRepository {
	return &NoteRepository{items: make(map[string]domain.FieldNote)}
}

func (r *NoteRepository) Create(item domain.FieldNote) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if item.ID == "" || item.ObservationID == "" || item.Author == "" {
		return errors.New("field note identity is incomplete")
	}
	if item.Text == "" {
		return errors.New("field note text is incomplete")
	}
	if _, exists := r.items[item.ID]; exists {
		return ErrNoteExists
	}
	r.items[item.ID] = item
	return nil
}

func (r *NoteRepository) ListForObservation(observationID string) []domain.FieldNote {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]domain.FieldNote, 0)
	for _, item := range r.items {
		if item.ObservationID == observationID {
			items = append(items, item)
		}
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].CreatedAt.Before(items[j].CreatedAt)
	})
	if len(items) == 0 {
		return []domain.FieldNote{}
	}
	return items
}

func (r *NoteRepository) CountForObservation(observationID string) int {
	return len(r.ListForObservation(observationID))
}
