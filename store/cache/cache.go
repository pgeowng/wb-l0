package cache

import (
	"context"

	"github.com/pgeowng/wb-l0/model"
	"github.com/pkg/errors"
)

type Cache struct {
	dict map[string]*model.Order
	keys []string
}

func New(ctx context.Context) (*Cache, error) {
	return &Cache{
		dict: make(map[string]*model.Order, 0),
		keys: make([]string, 0),
	}, nil
}

func (repo *Cache) Insert(ctx context.Context, order *model.Order) error {
	id := order.OrderUid
	_, ok := repo.dict[id]
	if ok {
		return errors.Errorf("cachemb: collision for id: %f", id)
	}

	repo.dict[id] = order
	repo.keys = append(repo.keys, id)

	return nil
}

func (repo *Cache) GetOrder(ctx context.Context, id string) (result *model.Order, err error) {
	data, ok := repo.dict[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return data, nil
}

func (repo *Cache) GetIds(ctx context.Context) (ids []string, err error) {
	result := make([]string, 0, len(repo.keys))
	for idx := range repo.keys {
		result = append(result, repo.keys[len(repo.keys)-idx-1])
	}

	return result, nil
}

func (repo *Cache) Recover(ctx context.Context, orders []*model.Order) error {
	repo.dict = make(map[string]*model.Order, len(orders))
	repo.keys = make([]string, 0, len(orders))

	for idx := range orders {
		id := orders[idx].OrderUid
		repo.dict[id] = orders[idx]
		repo.keys = append(repo.keys, id)
	}

	return nil
}
