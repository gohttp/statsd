package logger

import "github.com/statsd/client-interface"
import "net/http"
import "time"

// wrapper to capture status.
type wrapper struct {
	http.ResponseWriter
	written int
	status  int
}

// capture status.
func (w *wrapper) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// capture written bytes.
func (w *wrapper) Write(b []byte) (int, error) {
	n, err := w.ResponseWriter.Write(b)
	w.written += n
	return n, err
}

// New statsd middleware with the given statsd client.
func New(stats statsd.Client) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			start := time.Now()
			res := &wrapper{w, 0, 200}

			// request count
			stats.Incr("requests")

			// serve
			h.ServeHTTP(res, req)

			// status
			switch {
			case res.status >= 200 && res.status < 300:
				stats.Incr("response.ok")
			case res.status >= 400 && res.status < 500:
				stats.Incr("response.errors.client")
			case res.status >= 500:
				stats.Incr("response.errors.server")
			}

			// duration
			stats.Duration("response.duration", time.Since(start))

			// size
			stats.Histogram("response.size", res.written)
		})
	}
}
