package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
	"query-service/api/handlers"
	"query-service/config"
	"query-service/lib/pctx"
	"query-service/service/airTicket"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	service *airTicket.Service,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handlers.Ping(logger)).Methods(http.MethodGet)
	router.HandleFunc("/trips", handlers.Get(logger, service)).Methods(http.MethodGet)
	router.HandleFunc("/trips/{tripId}", handlers.GetById(logger, service)).Methods(http.MethodGet)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
