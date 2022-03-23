package pg

import (
	"github.com/pgeowng/wb-l0/model"
	"github.com/uptrace/bun"
)

type DBDelivery struct {
	bun.BaseModel `bun:"table:deliveries,alias:d"`
	*model.Delivery

	OrderUid string `bun:",pk"`
}

type DBPayment struct {
	bun.BaseModel `bun:"table:payments,alias:p"`
	*model.Payment

	OrderUid string `bun:",pk"`
}

type DBItem struct {
	bun.BaseModel `bun:"table:items,alias:i"`
	*model.Item

	OrderUid string `bun:",pk"`
}

type DBOrder struct {
	bun.BaseModel `bun:"table:orders,alias:o"`
	*model.Order

	OrderUid string `bun:",pk"`

	Delivery *DBDelivery `bun:"rel:has-one,join:order_uid=order_uid"`
	Payment  *DBPayment  `bun:"rel:has-one,join:order_uid=order_uid"`
	Items    []*DBItem   `bun:"rel:has-many,join:order_uid=order_uid"`
}
