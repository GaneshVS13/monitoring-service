package communication

import (
	"context"
	"net/http"
	"time"

	"github.com/monitoring-service/src/communication/rest"
	"github.com/monitoring-service/src/entity"
)

// Service is a service interface
type Service interface {
	MonitorService(ctx context.Context, url string) (entity.ServiceResponse, error)
}

// ServiceAPI is a service entity
type ServiceAPI struct {
	restSvc rest.RestService
}

// NewService returns new service instance
func NewService(s rest.RestService) Service {
	return &ServiceAPI{
		restSvc: s,
	}
}

// MonitorService makes the http request to external url
func (s *ServiceAPI) MonitorService(ctx context.Context, url string) (entity.ServiceResponse, error) {
	startTime := time.Now()

	_, statusCode, err := s.restSvc.Do(ctx, http.MethodGet, url, nil, nil)
	if err != nil {
		return entity.ServiceResponse{}, err
	}

	elapsed := time.Since(startTime).Milliseconds()

	return entity.ServiceResponse{
		StatusCode:   statusCode,
		ResponseTime: elapsed,
	}, nil
}
