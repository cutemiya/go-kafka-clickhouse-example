package airTicket

import (
	"context"
	"query-service/database/repo"
	"query-service/model"
)

type Service struct {
	repo *repo.Repository
}

func NewService(repo *repo.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (i *Service) Get(ctx context.Context) ([]model.OfferPrice, error) {
	return i.repo.Get(ctx)
}

func (i *Service) GetById(ctx context.Context, tripId int) ([]model.OfferPrice, error) {
	return i.repo.GetById(ctx, tripId)
}
