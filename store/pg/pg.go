package pg

import (
	"context"
	"database/sql"

	"github.com/pgeowng/wb-l0/model"
)

type DB struct {
	*sql.DB
}

func New(ctx context.Context) (*DB, error) {
	return &DB{}, nil
}

func (repo *DB) Insert(ctx context.Context, order *model.Order) error {
	return nil
}

func (repo *DB) GetAll(ctx context.Context) (result []*model.Order, err error) {
	return nil, nil
}
