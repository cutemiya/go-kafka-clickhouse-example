package handler

import (
	tripService "gateway-api/cluster/trip-service"
	"gateway-api/lib/responses"
	"gateway-api/model"
	"go.uber.org/zap"
	"net/http"
)

// CreateTrip
// @Description	create empty trip
// @Tags			docs
// @Produce		json
// @Success		200	{object}	model.TripResponse
// @Router			/trip [post]
func CreateTrip(logger *zap.SugaredLogger, service *tripService.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		createTripId, err := service.CreateTrip()
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
