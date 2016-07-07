package logging

import (
	"net/http"
)

// AddLoggingHandler will decorate the passed handler with a wrapper which
// will emit log messages whenever a request is received, in addition to
// calling the original handler.
func AddLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetLogger().Println("request:", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
