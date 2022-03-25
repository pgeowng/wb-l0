package store

import (
	"context"

	"github.com/pgeowng/wb-l0/model"
	"github.com/pgeowng/wb-l0/store/cache"
	"github.com/pgeowng/wb-l0/store/pg"
)

type InsertRepo interface {
	Insert(context.Context, *model.Order) error
}

type RestoreRepo interface {
	InsertRepo
	GetAll(context.Context) ([]*model.Order, error)
	Reset(context.Context) error
	Ensure(context.Context) error
}

type QueryRepo interface {
	InsertRepo
	GetOrder(context.Context, string) (*model.Order, error)
	GetIds(context.Context) ([]string, error)
	Recover(context.Context, []*model.Order) error
}

type Store struct {
	DB    RestoreRepo
	Cache QueryRepo
}

func New(ctx context.Context) (*Store, error) {
	db, err := pg.New(ctx)
	if err != nil {
		return nil, err
	}

	cache, err := cache.New(ctx)
	if err != nil {
		return nil, err
	}

	return &Store{
		DB:    db,
		Cache: cache,
	}, nil

}
