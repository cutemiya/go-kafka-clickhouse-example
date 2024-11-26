package trip

import (
	"context"
	"trip-service/database/repo"
)

type Service struct {
	repo *repo.Repository
}

func NewService(repository *repo.Repository) *Service {
	return &Service{
		repo: repository,
	}
}

func (i *Service) CreateTrip(ctx context.Context) (int, error) {
	return i.repo.CreateTrip(ctx)
}
