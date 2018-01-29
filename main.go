package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"github.com/CheerlessCloud/cexporter/collector"
	configPackage "github.com/CheerlessCloud/cexporter/config"
	_ "github.com/CheerlessCloud/cexporter/logger" // init logger

	log "github.com/CheerlessCloud/logrus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var config = configPackage.ConfigSingleton

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, os.Kill)

	server := &http.Server{
		Addr:    config.HTTPAddr,
		Handler: promhttp.HandlerFor(collector.Registry, promhttp.HandlerOpts{}),
	}

	go func() {
		collector.StartCollectingMetrics(config.FetchInterval, config.FetchTimeout)
	}()

	go func() {
		log.WithFields(log.Fields{"config": config}).Info("http server start")

		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	<-stop

	log.Info("Shutting down the server...")

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Error("Server shutting down emit error", err)
	} else {
		log.Info("Server gracefully stopped")
	}
}
