package logging

import (
	"fmt"
	"io"
	"net/http"

	"github.com/radanalyticsio/oshinko-rest/helpers/uuid"
)

// LogResponseWriter is a wrapper struct which allows us to retain the
// status code generated by ServeHTTP calls.
type logResponseWriter struct {
	writer   http.ResponseWriter
	status   int
	response *[]byte
}

func (w *logResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *logResponseWriter) Write(b []byte) (int, error) {
	w.response = &b
	return w.writer.Write(b)
}

func (w *logResponseWriter) WriteHeader(s int) {
	w.status = s
	w.writer.WriteHeader(s)
}

type logRequestReader struct {
	original io.ReadCloser
	Body     []byte
}

func (r *logRequestReader) Read(p []byte) (n int, err error) {
	n, err = r.original.Read(r.Body)
	copy(p, r.Body)
	return
}

func (r *logRequestReader) Close() error {
	return r.original.Close()
}

// AddLoggingHandler will decorate the passed handler with a wrapper which
// will emit log messages whenever a request is received, in addition to
// calling the original handler.
func AddLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := GetLogger()
		reqId, _ := uuid.Uuid()
		reqId = fmt.Sprintf("[req-id %s]", reqId)
		l.Println(reqId, r.Method, r.URL)
		lr := &logResponseWriter{w, 0, nil}
		br := &logRequestReader{r.Body, make([]byte, r.ContentLength)}
		r.Body = br
		next.ServeHTTP(lr, r)

		reqstr := fmt.Sprintf("%s", br.Body)
		Debug(reqId, "Request: ", reqstr)

		l.Println(reqId, lr.status)

		resstr := fmt.Sprintf("%s", *lr.response)
		Debug(reqId, "Response: ", resstr)
	})
}
