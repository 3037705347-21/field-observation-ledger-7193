package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"example.com/field-observation-ledger/internal/domain"
)

func writeJSON(writer http.ResponseWriter, status int, payload any) {
	writer.WriteHeader(status)
	_ = json.NewEncoder(writer).Encode(payload)
}

func writeSummary(writer http.ResponseWriter, status int, summary domain.Summary) {
	summary.EnsureMaps()
	writeJSON(writer, status, summary)
}

func writeError(writer http.ResponseWriter, status int, message string) {
	writeJSON(writer, status, map[string]string{"error": message})
}

func decodeJSON(request *http.Request, target any) error {
	if !acceptsJSON(request) {
		return errors.New("content type must be application/json")
	}
	defer request.Body.Close()
	decoder := json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}
