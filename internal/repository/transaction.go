package repository

import (
	"sync"

	"example.com/field-observation-ledger/internal/domain"
)

type Transaction struct {
	mu      sync.Mutex
	closed  bool
	changes []func()
}

func NewTransaction() *Transaction {
	return &Transaction{changes: make([]func(), 0)}
}

func (t *Transaction) Add(change func()) bool {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed || change == nil {
		return false
	}
	t.changes = append(t.changes, change)
	return true
}

func (t *Transaction) Commit() int {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.closed {
		return 0
	}
	changes := append([]func(){}, t.changes...)
	t.closed = true
	for _, change := range changes {
		change()
	}
	return len(changes)
}

func (t *Transaction) Rollback() {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.closed = true
	t.changes = nil
}

func CloneObservation(item domain.Observation) domain.Observation {
	item.Notes = string([]byte(item.Notes))
	return item
}

func CloneSpecies(item domain.Species) domain.Species {
	item.Tags = append([]string(nil), item.Tags...)
	return item
}
