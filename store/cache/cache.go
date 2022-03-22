package cache

import (
	"context"

	"github.com/pgeowng/wb-l0/model"
)

type Cache struct{}

func New(ctx context.Context) (*Cache, error) {
	return &Cache{}, nil
}

func (repo *Cache) Insert(context.Context, *model.Order) error {
	return nil
}

func (repo *Cache) GetOrder(ctx context.Context, id string) (order *model.Order, err error) {
	return nil, nil
}

func (repo *Cache) GetIds(ctx context.Context) (ids []string, err error) {

	return nil, nil
}

func (repo *Cache) Recover(ctx context.Context, orders []*model.Order) error {

	return nil
}
