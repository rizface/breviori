package observer

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
)

func RequestCounter() *prometheus.CounterVec {
	return prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "breviori",
		Name:      "breviori_http_request_total",
		Help:      "Count total http request",
	}, []string{"path", "code"})
}

func TimeResponseHistogram() *prometheus.HistogramVec {
	return prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "breviori",
		Name:      "breviori_response_time",
		Help:      "Observe the response time of breviori service",
		Buckets:   prometheus.DefBuckets,
	}, []string{"path", "status", "method"})
}

func RegisterMetrics(cs ...prometheus.Collector) error {
	for _, c := range cs {
		if err := prometheus.Register(c); err != nil {
			return fmt.Errorf("%s: %w", "failed register metrics: ", err)
		}
	}

	return nil
}
