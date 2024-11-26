package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"trip-service/model"

	_ "embed"
)

//go:embed query/insert.sql
var insertSql string

type Repository struct {
	logger *zap.SugaredLogger
	db     *sqlx.DB
}

func NewRepository(logger *zap.SugaredLogger, db *sqlx.DB) *Repository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

func (i *Repository) CreateTrip(ctx context.Context) (id int, err error) {
	err = i.db.QueryRowContext(ctx, insertSql, model.BookedStatus).Scan(&id)
	return id, err
}
