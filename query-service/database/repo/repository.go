package repo

import (
	"context"
	_ "embed"
	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"go.uber.org/zap"
	"query-service/model"
)

//go:embed query/getAll.sql
var getAllSql string

//go:embed query/get.sql
var getByIdSql string

type Repository struct {
	logger *zap.SugaredLogger
	db     driver.Conn
}

func NewRepository(logger *zap.SugaredLogger, db driver.Conn) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

func (i *Repository) GetById(ctx context.Context, tripId int) ([]model.OfferPrice, error) {
	var result []model.OfferPrice
	err := i.db.Select(ctx, &result, getByIdSql, tripId)

	return result, err
}

func (i *Repository) Get(ctx context.Context) ([]model.OfferPrice, error) {
	var result []model.OfferPrice
	err := i.db.Select(ctx, &result, getAllSql)

	return result, err
}
