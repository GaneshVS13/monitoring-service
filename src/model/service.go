package model

import (
	"context"
	"net/http"

	restClient "github.com/monitoring-service/src/communication"
	"github.com/monitoring-service/src/entity"
	"github.com/monitoring-service/src/web/metric"
)

// ServiceHandler handles the model operations
type ServiceHandler interface {
	Process(ctx context.Context)
}

type result struct {
	response entity.ServiceResponse
	url      string
	err      error
}

// ServiceModel represents the service model entity
type ServiceModel struct {
	urls            []string
	restService     restClient.Service
	metricPublisher metric.Publisher
}

// NewServiceModel returns the service model instance
func NewServiceModel(urls []string, metricPublisher metric.Publisher,
	restService restClient.Service) *ServiceModel {
	return &ServiceModel{
		urls:            urls,
		metricPublisher: metricPublisher,
		restService:     restService,
	}
}

// Process processes the service model request
func (s *ServiceModel) Process(ctx context.Context) {
	var (
		chResult = make(chan result, len(s.urls))
	)

	for _, url := range s.urls {
		go s.executeURL(ctx, url, chResult)
	}

	defer close(chResult)

	for range s.urls {
		output := <-chResult
		s.publishMetrics(output)
	}
}

func (s *ServiceModel) executeURL(ctx context.Context, url string, output chan<- result) {
	resp, err := s.restService.MonitorService(ctx, url)
	output <- result{
		url:      url,
		response: resp,
		err:      err,
	}
}

func (s *ServiceModel) publishMetrics(output result) {
	status := 0
	if output.response.StatusCode == http.StatusOK {
		status = 1
	}
	s.metricPublisher.PublishResponseStatus(output.url, status)
	s.metricPublisher.PublishResponseTime(output.url, output.response.ResponseTime)
}
