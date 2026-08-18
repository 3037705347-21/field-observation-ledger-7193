package service

import (
	"testing"

	"example.com/field-observation-ledger/internal/domain"
	"example.com/field-observation-ledger/internal/repository"
)

func TestSpeciesTagsAreIsolated(t *testing.T) {
	repository := repository.NewSpeciesRepository()
	service := NewSpeciesService(repository)
	_, err := service.Create(domain.Species{
		ID: "sp-isolated", CommonName: "Test species", ScientificName: "Testus isolatedus",
		Class: "bird", Tags: []string{"resident", "night"},
	})
	if err != nil {
		t.Fatal(err)
	}

	item, err := service.Get("sp-isolated")
	if err != nil {
		t.Fatal(err)
	}
	item.Tags[0] = "mutated"

	stored, err := service.Get("sp-isolated")
	if err != nil {
		t.Fatal(err)
	}
	if stored.Tags[0] != "resident" {
		t.Fatalf("repository state changed through returned tags: %#v", stored.Tags)
	}

	listed := service.List()
	for index := range listed {
		if listed[index].ID == "sp-isolated" {
			listed[index].Tags[1] = "listed-mutated"
		}
	}
	stored, err = service.Get("sp-isolated")
	if err != nil {
		t.Fatal(err)
	}
	if stored.Tags[1] != "night" {
		t.Fatalf("repository state changed through listed tags: %#v", stored.Tags)
	}
}
