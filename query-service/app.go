package main

import (
	"context"
	"net/http"
	"query-service/api"
	"query-service/config"
	"query-service/database"
	"query-service/database/repo"
	"query-service/lib/pctx"
	"query-service/service/airTicket"
	"time"

	"go.uber.org/zap"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	chCtx, cancelCtx := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelCtx()

	ch, err := database.NewClickhouse(chCtx, settings.Clickhouse)
	if err != nil {
		panic(err)
	}

	var (
		repository = repo.NewRepository(logger, ch)

		service = airTicket.NewService(repository)

		server = api.NewServer(ctxProvider, logger, settings, service)
	)

	return App{
		logger:   logger,
		settings: settings,
		server:   server,
	}
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
