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
	Recover(context.Context) error
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
		return errors.Wrap(err, "srv.create.db")
	}

	err = srv.store.Cache.Insert(ctx, order)
	if err != nil {
		return errors.Wrap(err, "srv.create.cache")
	}

	return nil
}

func (srv *OrderServiceImpl) GetOrder(ctx context.Context, id string) (result *model.Order, err error) {
	result, err = srv.store.Cache.GetOrder(ctx, id)
	err = errors.Wrap(err, "srv")
	return
}

func (srv *OrderServiceImpl) GetIds(ctx context.Context) (ids []string, err error) {
	ids, err = srv.store.Cache.GetIds(ctx)
	err = errors.Wrap(err, "srv")
	return
}

func (srv *OrderServiceImpl) Recover(ctx context.Context) error {
	data, err := srv.store.DB.GetAll(ctx)
	if err != nil {
		return errors.Wrap(err, "recover.db")
	}

	err = srv.store.Cache.Recover(ctx, data)
	if err != nil {
		return errors.Wrap(err, "recover")
	}

	return nil
}
