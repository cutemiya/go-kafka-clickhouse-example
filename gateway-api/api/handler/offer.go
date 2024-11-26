package handler

import (
	"encoding/json"
	"gateway-api/lib/responses"
	"gateway-api/model"
	offerService "gateway-api/service/offer-service"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// Book
//
//	@Description	Book trip
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			request	body	model.Offer	true	"param"
//	@Success		200
//	@Router			/trip-book [post]
func Book(logger *zap.SugaredLogger, offerServiceImpl *offerService.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request model.Offer

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			responses.SendInternalServerError(w)
			return
		}

		tripId, err := offerServiceImpl.BookTrip(r.Context(), request)
		if err != nil {
			logger.Debugf("Book error: %v", err)
			responses.SendInternalServerError(w)
			return
		}

		res := model.TripResponse{
			Id: tripId,
		}

		responses.SendResponse(w, http.StatusOK, res.Marshal())
	}
}

// UpdateOffer
//
//	@Description	push new offer to kafka topic
//	@Tags			docs
//	@Accept			json
//	@Produce		json
//	@Param			id	path	int	true	"id"
//	@Param			request	body	model.Offer	true	"param"
//	@Success		200
//	@Router			/trip/{tripId}/offer [post]
func UpdateOffer(logger *zap.SugaredLogger, offerServiceImpl *offerService.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request model.Offer

		id, _ := strconv.Atoi(mux.Vars(r)["id"])
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&request)
		if err != nil {
			logger.Debugf("Decode request error: %s", err)
			responses.SendInternalServerError(w)
			return
		}

		request.TripId = id

		err = offerServiceImpl.UpdateOffer(r.Context(), request)
		if err != nil {
			logger.Debugf("Book error: %v", err)
			responses.SendInternalServerError(w)
			return
		}

		res := model.TripResponse{
			Id: id,
		}

		responses.SendResponse(w, http.StatusOK, res.Marshal())
	}
}
