package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"trip-service/api/handler"
	"trip-service/config"
	"trip-service/lib/pctx"
	"trip-service/service/trip"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	service *trip.Service,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/trip", handler.CreateTrip(logger, service)).Methods(http.MethodPost)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
