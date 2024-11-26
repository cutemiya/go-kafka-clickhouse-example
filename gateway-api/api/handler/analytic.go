package handler

import (
	"gateway-api/cluster/queryService"
	"gateway-api/lib/responses"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Get
//
//	@Description	get all trips
//	@Tags			analytics
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.OfferPrice
//	@Router			/trips [get]
func Get(logger *zap.SugaredLogger, service *queryService.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := service.GetAll()
		if err != nil {
			logger.Fatalf(err.Error())
			responses.SendInternalServerError(w)
		}

		responses.SendResponse(w, http.StatusOK, res)
	}
}

// GetById
//
//	@Description	get trip by id
//	@Tags			analytics
//	@Accept			json
//	@Produce		json
//	@Param			tripId	path		int	true	"tripId"
//	@Success		200	{object}	model.OfferPrice
//	@Router			/trips/{tripId} [get]
func GetById(logger *zap.SugaredLogger, service *queryService.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["tripId"])

		res, err := service.GetById(id)
		if err != nil {
			logger.Debugf(err.Error())
			responses.SendInternalServerError(w)
			return
		}

		responses.SendResponse(w, http.StatusOK, res)
	}
}
