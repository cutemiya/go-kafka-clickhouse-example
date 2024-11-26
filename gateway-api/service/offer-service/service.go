package offerService

import (
	"context"
	"encoding/json"
	"errors"
	tripService "gateway-api/cluster/trip-service"
	"gateway-api/model"
	"github.com/segmentio/kafka-go"
)

var KafkaProduceError = errors.New("не удалось запушить в кафку")

type Service struct {
	tripClient *tripService.Service
	writer     *kafka.Writer
}

func NewService(client *tripService.Service, writer *kafka.Writer) *Service {
	return &Service{
		tripClient: client,
		writer:     writer,
	}
}

func (i *Service) BookTrip(ctx context.Context, offer model.Offer) (tripId int, err error) {
	tripId, err = i.tripClient.CreateTrip()
	if err != nil {
		return 0, err
	}
	offer.TripId = tripId

	jsonModel, err := json.Marshal(offer)
	if err != nil {
		return 0, err
	}

	err = i.writer.WriteMessages(ctx, kafka.Message{
		Value: jsonModel,
	})
	if err != nil {
		return tripId, KafkaProduceError
	}

	return tripId, nil
}

func (i *Service) UpdateOffer(ctx context.Context, offer model.Offer) error {
	jsonModel, err := json.Marshal(offer)
	if err != nil {
		return err
	}

	err = i.writer.WriteMessages(ctx, kafka.Message{
		Value: jsonModel,
	})
	if err != nil {
		return KafkaProduceError
	}

	return nil
}
