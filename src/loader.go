package main

import (
	"context"
	"log"
	"net/http"

	service "github.com/monitoring-service/src/communication"
	"github.com/monitoring-service/src/communication/rest"
	"github.com/monitoring-service/src/config"
	"github.com/monitoring-service/src/model"
	"github.com/monitoring-service/src/web"
	"github.com/monitoring-service/src/web/metric"
	"github.com/monitoring-service/src/web/middleware"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type StopFunc func()

// loadApplication loads the application
func loadApplication(ctx context.Context, configFilePath string) (StopFunc, error) {
	// load config
	cfg, err := config.Load(configFilePath)
	if err != nil {
		return nil, err
	}

	// load communication clients
	restService := rest.NewService()
	serviceClient := service.NewService(restService)

	// load metric publisher
	metricPublisher, promRegistry := metric.NewMetricPublisher()

	// load models
	serviceModel := model.NewServiceModel(cfg.MonitoringURLs, metricPublisher, serviceClient)

	// Init http server and handler instances
	router := mux.NewRouter().StrictSlash(true)
	web.NewHandler(cfg.ServiceConfig, router, promRegistry)

	server := &http.Server{
		Addr:    cfg.ServiceConfig.ListenURL,
		Handler: middleware.NewMiddleware(handlers.CORS()(router), serviceModel),
	}

	initAndStartServer(ctx, server)

	return func() {
	}, nil
}

func initAndStartServer(ctx context.Context, server *http.Server) {
	go func() {
		log.Println(ctx, "starting HTTP listener...")
		err := server.ListenAndServe()
		if err != nil {
			log.Println(ctx, "server has been stopped.", err)
		}
	}()
}
