package service

import (
	"errors"
	"fmt"
	"strings"

	"example.com/field-observation-ledger/internal/clock"
	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

var ErrNoteObservation = errors.New("field note observation does not exist")
var ErrNoteAuthor = errors.New("field note author is required")
var ErrNoteText = errors.New("field note text is required")

type NoteService struct {
	observations *ObservationService
	repository   *repository.NoteRepository
	clock        clock.Clock
}

func NewNoteService(
	observations *ObservationService,
	repo *repository.NoteRepository,
	timeSource clock.Clock,
) *NoteService {
	return &NoteService{observations: observations, repository: repo, clock: timeSource}
}

func (s *NoteService) Add(observationID, author, text string, tags []string) (domain.FieldNote, error) {
	if _, err := s.observations.Get(observationID); err != nil {
		return domain.FieldNote{}, ErrNoteObservation
	}
	author = strings.TrimSpace(author)
	text = strings.TrimSpace(text)
	if author == "" {
		return domain.FieldNote{}, ErrNoteAuthor
	}
	if text == "" {
		return domain.FieldNote{}, ErrNoteText
	}
	note := domain.FieldNote{
		ID:            fmt.Sprintf("note-%s-%d", observationID, s.clock.Now().UnixNano()),
		ObservationID: observationID,
		Author:        author,
		Text:          text,
		Tags:          cleanTags(tags),
		CreatedAt:     s.clock.Now(),
	}
	if err := s.repository.Create(note); err != nil {
		return domain.FieldNote{}, err
	}
	return note.Clone(), nil
}

func (s *NoteService) List(observationID string) []domain.FieldNote {
	return cloneNotes(s.repository.ListForObservation(observationID))
}

func (s *NoteService) Count(observationID string) int {
	return s.repository.CountForObservation(observationID)
}

func cleanTags(tags []string) []string {
	result := make([]string, 0, len(tags))
	seen := make(map[string]bool)
	for _, tag := range tags {
		value := strings.ToLower(strings.TrimSpace(tag))
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}

func cloneNotes(items []domain.FieldNote) []domain.FieldNote {
	clones := make([]domain.FieldNote, len(items))
	for i, item := range items {
		clones[i] = item.Clone()
	}
	return clones
}
