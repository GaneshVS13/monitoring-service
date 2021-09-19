package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

const (
	configFile = "./monitoring_service_cfg.json"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	stop, err := loadApplication(ctx, configFile)
	if err != nil {
		panic(err)
	}

	defer stop()
	defer cancel()

	s := make(chan os.Signal, 1)
	signal.Notify(s,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	<-s
}
