package cache

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/pgeowng/wb-l0/model"
	"github.com/pkg/errors"
)

type Cache struct {
	dict map[string][]byte
	keys []string
}

func New(ctx context.Context) (*Cache, error) {
	return &Cache{
		dict: make(map[string][]byte, 0),
		keys: make([]string, 0),
	}, nil
}

func (repo *Cache) Insert(ctx context.Context, order *model.Order) error {
	id := *order.OrderUid
	_, ok := repo.dict[id]
	if ok {
		return errors.Errorf("cachemb: collision for id: %f", id)
	}

	data, err := json.Marshal(order)
	if err != nil {
		return errors.Wrap(err, "cachemb")
	}

	repo.dict[id] = data
	repo.keys = append(repo.keys, id)
	sort.Strings(repo.keys)

	return nil
}

func (repo *Cache) GetOrder(ctx context.Context, id string) (result []byte, err error) {

	data, ok := repo.dict[id]
	if !ok {
		return nil, errors.New("not found")
	}

	return data, nil
}

func (repo *Cache) GetIds(ctx context.Context) (ids []string, err error) {
	return repo.keys, nil
}

func (repo *Cache) Recover(ctx context.Context, orders []*model.Order) error {

	repo.dict = make(map[string][]byte, len(orders))
	repo.keys = make([]string, 0, len(orders))

	for idx := range orders {
		id := *orders[idx].OrderUid
		data, err := json.Marshal(orders[idx])
		if err != nil {
			return errors.Wrap(err, "cachemb")
		}
		repo.dict[id] = data
		repo.keys = append(repo.keys, id)
	}

	sort.Strings(repo.keys)

	return nil
}
