package api

import (
	"context"
	"fmt"
	"gateway-api/api/handler"
	"gateway-api/cluster/queryService"
	tripService "gateway-api/cluster/trip-service"
	"gateway-api/config"
	"gateway-api/lib/pctx"
	offerService "gateway-api/service/offer-service"
	"github.com/gorilla/mux"
	//"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"net"
	"net/http"
	//"trip-service/service/trip"

	httpSwagger "github.com/swaggo/http-swagger"
)

func NewServer(
	ctxProvider pctx.DefaultProvider,
	logger *zap.SugaredLogger,
	settings config.Settings,
	tripClusterClient *tripService.Service,
	offerServiceImpl *offerService.Service,
	queryClusterClient *queryService.Service,
) *http.Server {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handler.Ping(logger)).Methods(http.MethodGet)

	router.HandleFunc("/trip", handler.CreateTrip(logger, tripClusterClient)).Methods(http.MethodPost)
	router.HandleFunc("/trip-book", handler.Book(logger, offerServiceImpl)).Methods(http.MethodPost)
	router.HandleFunc("/trip/{tripId}/offer", handler.UpdateOffer(logger, offerServiceImpl)).Methods(http.MethodPost)

	router.HandleFunc("/trips", handler.Get(logger, queryClusterClient)).Methods(http.MethodGet)
	router.HandleFunc("/trips/{tripId}", handler.GetById(logger, queryClusterClient)).Methods(http.MethodGet)
	//router.HandleFunc("/ticket", handler.Add(logger, insertProducer)).Methods(http.MethodPost)
	//router.HandleFunc("/ticket", handler.Update(logger, updateProducer)).Methods(http.MethodPut)
	//router.HandleFunc("/ticket/{id}", handler.Delete(logger, deleteProducer)).Methods(http.MethodDelete)
	//router.HandleFunc("/ticket-wo-cache/{id}", handler.GetByIdWOCache(logger, service)).Methods(http.MethodGet)

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return &http.Server{
		Addr: fmt.Sprintf(":%d", settings.Port),
		BaseContext: func(listener net.Listener) context.Context {
			return ctxProvider()
		},
		Handler: router,
	}
}
