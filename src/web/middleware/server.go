package middleware

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/monitoring-service/src/entity"
	"github.com/monitoring-service/src/model"
)

// MetricMiddleware represents the entity for metric middleware
type MetricMiddleware struct {
	serviceModel model.ServiceHandler
}

// NewMetricMiddleware returns metric middleware instance
func NewMetricMiddleware(serviceModel model.ServiceHandler) *MetricMiddleware {
	return &MetricMiddleware{
		serviceModel: serviceModel,
	}
}

// ServeHTTP serves the http request
func (m MetricMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	transactionID := r.Header.Get(entity.TransactionIDKey)
	if transactionID == "" {
		transactionID = uuid.New().String()
	}

	ctx := r.Context()
	ctx = entity.SetTransactionContextValue(ctx, transactionID)

	m.serviceModel.Process(ctx)

	next(w, r.WithContext(ctx))
}
