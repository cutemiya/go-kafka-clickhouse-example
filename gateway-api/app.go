package main

import (
	"context"
	"gateway-api/api"
	"gateway-api/cluster/queryService"
	tripService "gateway-api/cluster/trip-service"
	"gateway-api/config"
	kafkaSettings "gateway-api/kafka"
	"gateway-api/lib/pctx"
	offerService "gateway-api/service/offer-service"
	"github.com/segmentio/kafka-go"
	"net/http"

	"go.uber.org/zap"
)

type Producers struct {
	InsertProducer *kafka.Writer
}

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func InitProducers(settings config.Settings) Producers {
	var addr = settings.KafkaSettings.Brokers[0]

	var producers Producers

	producers.InsertProducer = kafkaSettings.NewWriter(addr, settings.KafkaSettings.Topics.InsertTopic)

	return producers
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	var (
		httpClient         = &http.Client{}
		queryClusterClient = queryService.NewService(httpClient, settings.QueryService)
		tripClusterClient  = tripService.NewService(httpClient, settings.TripService)
		producers          = InitProducers(settings)
		offerServiceImpl   = offerService.NewService(tripClusterClient, producers.InsertProducer)
		server             = api.NewServer(
			ctxProvider,
			logger,
			settings,
			tripClusterClient,
			offerServiceImpl,
			queryClusterClient,
		)
	)

	app := App{
		logger:   logger,
		settings: settings,
		server:   server,
	}

	return app
}

func (a App) Run() {
	go func() {
		_ = a.server.ListenAndServe()
	}()
	a.logger.Debugf("HTTP handler started on %d", a.settings.Port)
}

func (a App) Stop(ctx context.Context) {
	_ = a.server.Shutdown(ctx)
	a.logger.Debugf("HTTP handler stopped")
}
