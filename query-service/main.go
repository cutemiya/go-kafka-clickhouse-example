package main

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"query-service/config"
	_ "query-service/docs"
	"syscall"
	"time"
)

//	@title			Swagger of API
//	@version		1.0
//	@description	Swagger for SOA

// @host		localhost:80
// @BasePath	/
func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	var (
		loggerSugar = logger.Sugar()
		settings    = config.Load(loggerSugar)
	)

	mainCtx := context.Background()
	notifyCtx, cancelFunc := signal.NotifyContext(mainCtx, os.Interrupt, syscall.SIGTERM)
	defer cancelFunc()

	app := NewApp(func() context.Context {
		ctx, _ := context.WithCancel(notifyCtx)
		return ctx
	}, loggerSugar, settings)
	app.Run()

	select {
	case <-notifyCtx.Done():
	}
	loggerSugar.Debug("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	app.Stop(ctx)
	loggerSugar.Debug("Successful stopped")
}
