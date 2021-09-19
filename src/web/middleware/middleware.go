package middleware

import (
	"net/http"

	"github.com/monitoring-service/src/model"

	"github.com/urfave/negroni"
)

// NewMiddleware new http handler
func NewMiddleware(router http.Handler,
	serviceModel model.ServiceHandler) http.Handler {
	middlewareManager := negroni.New()
	middlewareManager.Use(negroni.NewRecovery())
	middlewareManager.Use(NewMetricMiddleware(serviceModel))
	middlewareManager.UseHandler(router)

	return middlewareManager
}
