package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

// MetricPublisher represents the metric publisher entity
type MetricPublisher struct {
	responseStatus   *prometheus.GaugeVec
	responseDuration *prometheus.GaugeVec
}

// Publisher handles the metric publisher operations
type Publisher interface {
	PublishResponseStatus(label string, val int)
	PublishResponseTime(label string, duration int64)
}
