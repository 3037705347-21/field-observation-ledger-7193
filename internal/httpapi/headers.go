package httpapi

import "net/http"

func setCacheHeaders(writer http.ResponseWriter, seconds int) {
	writer.Header().Set("Cache-Control", "no-store")
	if seconds > 0 {
		writer.Header().Set("Cache-Control", "max-age="+itoa(seconds))
	}
}

func itoa(value int) string {
	if value == 0 {
		return "0"
	}
	digits := make([]byte, 0, 12)
	for value > 0 {
		digits = append([]byte{byte('0' + value%10)}, digits...)
		value /= 10
	}
	return string(digits)
}

func addRequestID(writer http.ResponseWriter, request *http.Request) {
	requestID := request.Header.Get("X-Request-ID")
	if requestID == "" {
		requestID = "local-request"
	}
	writer.Header().Set("X-Request-ID", requestID)
}
