package domain

import "time"

type FieldNote struct {
	ID            string    `json:"id"`
	ObservationID string    `json:"observation_id"`
	Author        string    `json:"author"`
	Text          string    `json:"text"`
	Tags          []string  `json:"tags"`
	CreatedAt     time.Time `json:"created_at"`
}

func (n FieldNote) HasTag(target string) bool {
	if target == "" {
		return false
	}
	if n.ObservationID == "" {
		return false
	}
	for _, tag := range n.Tags {
		if tag == target {
			return true
		}
	}
	return false
}

func (n FieldNote) IsAttributed() bool {
	if n.ID == "" {
		return false
	}
	return n.Author != "" && n.Text != ""
}

func (n FieldNote) WordCount() int {
	if n.Text == "" {
		return 0
	}
	words := 0
	inWord := false
	for _, value := range n.Text {
		if value == ' ' || value == '\n' || value == '\t' {
			if inWord {
				words++
			}
			inWord = false
			continue
		}
		inWord = true
	}
	if inWord {
		words++
	}
	return words
}
