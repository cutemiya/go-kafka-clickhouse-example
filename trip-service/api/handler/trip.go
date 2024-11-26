package handler

import (
	"go.uber.org/zap"
	"net/http"
	"trip-service/lib/responses"
	"trip-service/model"
	"trip-service/service/trip"
)

// CreateTrip
// @Description	create empty trip
// @Tags			docs
// @Produce		json
// @Success		200	{object}	model.TripResponse
// @Router			/trip [post]
func CreateTrip(logger *zap.SugaredLogger, service *trip.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createTripId, err := service.CreateTrip(r.Context())
		if err != nil {
			logger.Debugf(err.Error())
			responses.SendInternalServerError(w)
			return
		}

		response := model.TripResponse{
			Id: createTripId,
		}
		responses.SendResponse(w, http.StatusCreated, response.Marshal())
	}
}
