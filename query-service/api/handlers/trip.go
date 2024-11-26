package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"query-service/lib/responses"
	"query-service/service/airTicket"
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
func Get(logger *zap.SugaredLogger, service *airTicket.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := service.Get(r.Context())
		if err != nil {
			logger.Fatalf(err.Error())
			responses.SendInternalServerError(w)
		}

		responseJson, _ := json.Marshal(res)
		responses.SendResponse(w, http.StatusOK, responseJson)
	}
}

// GetById
//
//	@Description	get trip by id
//	@Tags			analytics
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"tripId"
//	@Success		200	{object}	model.OfferPrice
//	@Router			/trips/{tripId} [get]
func GetById(logger *zap.SugaredLogger, service *airTicket.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(mux.Vars(r)["tripId"])

		res, err := service.GetById(r.Context(), id)
		if err != nil {
			logger.Debugf(err.Error())
			responses.SendInternalServerError(w)
			return
		}

		responseJson, _ := json.Marshal(res)
		responses.SendResponse(w, http.StatusOK, responseJson)
	}
}
