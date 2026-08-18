package httpapi

import (
	"log"
	"net/http"
	"time"
)

func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		setCacheHeaders(writer, 0)
		addRequestID(writer, request)
		next.ServeHTTP(writer, request)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		started := time.Now()
		next.ServeHTTP(writer, request)
		log.Printf("%s %s %s", request.Method, request.URL.Path, time.Since(started).Round(time.Millisecond))
	})
}
