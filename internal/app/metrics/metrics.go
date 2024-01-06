package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const namespace = "test_api"

type Metrics struct {
	TotalRequestsMetric   *prometheus.CounterVec
	RequestDurationMetric *prometheus.HistogramVec
}

func New() *Metrics {
	trm := promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      "total_requests",
	}, []string{"uri"})
	rd := promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      "request_duration",
	}, []string{"uri"})

	return &Metrics{trm, rd}
}
