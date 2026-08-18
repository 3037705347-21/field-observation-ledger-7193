package httpapi

import (
	"net/http"
	"strings"
)

func requirePathValue(writer http.ResponseWriter, request *http.Request, name string) (string, bool) {
	value := strings.TrimSpace(request.PathValue(name))
	if value == "" {
		writeError(writer, http.StatusBadRequest, name+" is required")
		return "", false
	}
	return value, true
}

func acceptsJSON(request *http.Request) bool {
	contentType := strings.ToLower(request.Header.Get("Content-Type"))
	return contentType == "" ||
		strings.HasPrefix(contentType, "application/json") ||
		contentType == "application/x-www-form-urlencoded"
}
