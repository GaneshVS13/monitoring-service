package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	guageLabel = "url"
)

// NewMetricPublisher returns the metric publisher instance
func NewMetricPublisher() (*MetricPublisher, *prometheus.Registry) {
	responseStatus := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "sample_external_url_up",
		Help: "Status of HTTP response",
	}, []string{guageLabel})

	responseTime := promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "sample_external_url_response_ms",
		Help: "Duration of HTTP request in miliseconds",
	}, []string{guageLabel})

	publisher := &MetricPublisher{
		responseStatus:   responseStatus,
		responseDuration: responseTime,
	}

	return publisher, publisher.register()
}

func (mp *MetricPublisher) register() *prometheus.Registry {
	registry := prometheus.NewRegistry()
	registry.MustRegister(mp.responseStatus)
	registry.MustRegister(mp.responseDuration)
	return registry
}

// PublishResponseStatus publishes the response status metric
func (mp *MetricPublisher) PublishResponseStatus(label string, val int) {
	mp.responseStatus.WithLabelValues(label).Set(float64(val))
}

// PublishResponseTime publishes the response time metric
func (mp *MetricPublisher) PublishResponseTime(label string, val int64) {
	mp.responseDuration.WithLabelValues(label).Set(float64(val))
}
