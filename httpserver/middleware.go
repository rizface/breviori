package httpserver

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Recorder struct {
	HttpRequestCounter   *prometheus.CounterVec
	ResponseTimeObserver *prometheus.HistogramVec
}

type CustomResponseWriter struct {
	http.ResponseWriter
	StatusCode int
}

func (c *CustomResponseWriter) WriteHeader(statusCode int) {
	c.ResponseWriter.WriteHeader(statusCode)
	c.StatusCode = statusCode
}

func MetricsRecorder(rec Recorder) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			now := time.Now()
			cw := &CustomResponseWriter{ResponseWriter: w}

			next.ServeHTTP(cw, r)

			time.Sleep(
				time.Duration(rand.Intn(5)) * time.Second,
			)

			statusCode := strconv.Itoa(cw.StatusCode)

			rec.HttpRequestCounter.WithLabelValues(r.URL.Path, statusCode).Inc()
			rec.ResponseTimeObserver.WithLabelValues(r.URL.Path, r.Method, statusCode).Observe(float64(time.Since(now).Seconds()))
		})
	}
}
