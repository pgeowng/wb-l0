package pg

import (
	"context"
	"database/sql"

	"github.com/pgeowng/wb-l0/model"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

func (db *DB) GetAll(ctx context.Context) (result []*model.Order, err error) {

	return nil, nil
}

func cancelTx(tx bun.Tx, err error) error {
	e := tx.Rollback()
	if e != nil {
		return errors.Wrapf(e, "pg.tx.rollback(%s)", err)
	}

	return errors.Wrap(err, "pg.tx")
}

func (db *DB) Insert(ctx context.Context, order *model.Order) error {
	orderUid := order.OrderUid

	tx, err := db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return errors.Wrap(err, "pg.tx")
	}

	delivery := &DBDelivery{
		Delivery: &order.Delivery,
		OrderUid: orderUid,
	}
	_, err = tx.NewInsert().Model(delivery).Exec(ctx)
	if err != nil {
		return cancelTx(tx, err)
	}

	payment := &DBPayment{
		Payment:  &order.Payment,
		OrderUid: orderUid,
	}
	_, err = tx.NewInsert().Model(payment).Exec(ctx)
	if err != nil {
		return cancelTx(tx, err)
	}

	if len(order.Items) > 0 {
		items := []DBItem{}
		for idx := range order.Items {
			items = append(items, DBItem{
				Item:     &(order.Items[idx]),
				OrderUid: orderUid,
			})
		}
		_, err = tx.NewInsert().Model(&items).Exec(ctx)
		if err != nil {
			return cancelTx(tx, err)
		}
	}

	dbOrder := &DBOrder{
		Order:    order,
		OrderUid: orderUid,
	}
	_, err = tx.NewInsert().Model(dbOrder).Exec(ctx)
	if err != nil {
		return cancelTx(tx, err)
	}

	err = tx.Commit()
	if err != nil {
		return errors.Wrap(err, "pg.tx")
	}

	return nil
}
