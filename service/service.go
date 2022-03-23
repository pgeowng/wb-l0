package service

import (
	"context"

	"github.com/pgeowng/wb-l0/model"
	"github.com/pgeowng/wb-l0/store"
	"github.com/pkg/errors"
)

type OrderService interface {
	Create(context.Context, *model.Order) error
	GetOrder(context.Context, string) (*model.Order, error)
	GetIds(context.Context) ([]string, error)
}

type OrderServiceImpl struct {
	store *store.Store
}

func New(ctx context.Context, store *store.Store) OrderService {
	return &OrderServiceImpl{store: store}
}

func (srv *OrderServiceImpl) Create(ctx context.Context, order *model.Order) error {
	err := srv.store.DB.Insert(ctx, order)
	if err != nil {
		return errors.Wrap(err, "srv.create")
	}

	return nil
}

func (srv *OrderServiceImpl) GetOrder(ctx context.Context, id string) (order *model.Order, err error) {
	return nil, nil
}

func (srv *OrderServiceImpl) GetIds(ctx context.Context) (ids []string, err error) {
	return nil, nil
}
