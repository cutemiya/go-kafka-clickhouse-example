package main

import (
	"context"
	"net/http"
	"trip-service/api"
	"trip-service/config"
	"trip-service/database"
	"trip-service/database/repo"
	"trip-service/lib/pctx"
	"trip-service/service/trip"

	"go.uber.org/zap"
)

type App struct {
	logger   *zap.SugaredLogger
	settings config.Settings
	server   *http.Server
}

func NewApp(ctxProvider pctx.DefaultProvider, logger *zap.SugaredLogger, settings config.Settings) App {
	pgDb, err := database.NewPgx(settings.Postgres)
	if err != nil {
		panic(err)
	}
	err = database.UpMigrations(pgDb)
	if err != nil {
		panic(err)
	}

	var (
		repository = repo.NewRepository(logger, pgDb)
		service    = trip.NewService(repository)
		server     = api.NewServer(
			ctxProvider,
			logger,
			settings,
			service,
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
