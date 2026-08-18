# Field Observation Ledger

Field Observation Ledger is an offline-friendly Go HTTP service for ecological field teams. It keeps observation records connected to species and field sites, then supports review decisions and summary queries without an external database.

## Users and modules

- Observers submit draft observations.
- Catalogers maintain species and site references.
- Reviewers approve or reject observations.
- Research coordinators query filtered observations and summaries.

The `internal/domain` package defines business records, `internal/repository` provides synchronized in-memory storage, `internal/service` enforces rules, and `internal/httpapi` exposes the HTTP endpoints.

## Run

```text
go run ./cmd/observer
```

The default server listens on `127.0.0.1:18080`. Set `OBSERVE_ADDR` to override the address.

## Build and test

```text
go build ./...
go test ./...
```

## API examples

```text
GET  /health
GET  /api/v1/sites
GET  /api/v1/species
GET  /api/v1/observations?site_id=site-alpine&status=approved
POST /api/v1/observations
POST /api/v1/reviews/{observation-id}/approve
POST /api/v1/reviews/{observation-id}/reject
GET  /api/v1/summary?site_id=site-alpine
```

Observation creation expects JSON with `id`, `site_id`, `species_id`, `observed_at`, `count`, and optional `notes`. Review creation expects `reviewer` and `comment`.
