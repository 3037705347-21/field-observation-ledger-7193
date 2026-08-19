package httpapi

import (
	"net/http"
	"strconv"
	"strings"
)
type pageOptions struct { Offset int; Limit int }
func parsePage(request *http.Request, defaultLimit int) pageOptions {
	if defaultLimit <= 0 { defaultLimit = 50 }
	options := pageOptions{Limit: defaultLimit}
	if raw := strings.TrimSpace(request.URL.Query().Get("offset")); raw != "" {
		if value, err := strconv.Atoi(raw); err == nil && value >= 0 { options.Offset = value }
	}
	if raw := strings.TrimSpace(request.URL.Query().Get("limit")); raw != "" {
		if value, err := strconv.Atoi(raw); err == nil && value > 0 && value <= 200 { options.Limit = value }
	}
	return options
}
func paginate[T any](items []T, options pageOptions) []T {
	if options.Offset >= len(items) { return []T{} }
	end := options.Offset + options.Limit
	if end > len(items) { end = len(items) }
	return items[options.Offset:end]
}
