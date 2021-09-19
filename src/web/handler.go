package web

import (
	"github.com/monitoring-service/src/config"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Handler is a web request handler
type Handler struct {
	cfg      config.ServiceConfig
	registry *prometheus.Registry
}

// NewHandler creates a handler instance
func NewHandler(cfg config.ServiceConfig,
	router *mux.Router,
	registry *prometheus.Registry) {
	handler := Handler{
		cfg:      cfg,
		registry: registry,
	}

	handler.routesBuilder(router)
}

// routesBuilder creates web routes
func (h Handler) routesBuilder(router *mux.Router) {
	//api := router.PathPrefix(h.cfg.URLPrefix).Subrouter()
	//apiV1 := api.PathPrefix(h.cfg.APIVersion).Subrouter()

	router.Handle("/metrics", promhttp.HandlerFor(h.registry, promhttp.HandlerOpts{}))
}
