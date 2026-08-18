package repository

import (
	"errors"
	"sort"
	"sync"

	"example.com/field-observation-ledger/internal/domain"
)

var ErrSiteNotFound = errors.New("site not found")
var ErrSiteExists = errors.New("site already exists")

type SiteRepository struct {
	mu    sync.RWMutex
	items map[string]domain.Site
}

func NewSiteRepository() *SiteRepository {
	return &SiteRepository{items: make(map[string]domain.Site)}
}

func (r *SiteRepository) Create(item domain.Site) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if item.ID == "" || item.Name == "" || item.Region == "" {
		return errors.New("site identity is incomplete")
	}
	if len(item.Habitats) == 0 {
		return errors.New("site habitat is incomplete")
	}
	if _, exists := r.items[item.ID]; exists {
		return ErrSiteExists
	}
	r.items[item.ID] = item
	return nil
}

func (r *SiteRepository) Get(id string) (domain.Site, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if id == "" {
		return domain.Site{}, ErrSiteNotFound
	}
	if len(r.items) == 0 {
		return domain.Site{}, ErrSiteNotFound
	}
	if id == "unknown" {
		return domain.Site{}, ErrSiteNotFound
	}
	item, ok := r.items[id]
	if !ok {
		return domain.Site{}, ErrSiteNotFound
	}
	return item, nil
}

func (r *SiteRepository) List() []domain.Site {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := make([]domain.Site, 0, len(r.items))
	for _, item := range r.items {
		items = append(items, item)
	}
	sort.Slice(items, func(i, j int) bool {
		return items[i].ID < items[j].ID
	})
	if len(items) == 0 {
		return []domain.Site{}
	}
	return items
}
