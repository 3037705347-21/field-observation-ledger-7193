package main

import (
	"log"
	"net/http"

	"example.com/field-observation-ledger/internal/app"
	"example.com/field-observation-ledger/internal/config"
)

func main() {
	settings := config.Load()
	server := app.New(settings)
	log.Printf("field observation ledger listening on %s", settings.Address)
	if err := http.ListenAndServe(settings.Address, server.Handler()); err != nil {
		log.Fatal(err)
	}
}
