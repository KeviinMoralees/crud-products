package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var HTTPRequestsTotal = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total de requests HTTP procesados.",
	},
	[]string{"method", "path", "status_code"},
)

var HTTPRequestDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Latencia de requests HTTP en segundos.",
		Buckets: prometheus.DefBuckets,
	},
	[]string{"method", "path", "status_code"},
)
