package pg

import (
	"context"
	"database/sql"

	"github.com/pgeowng/wb-l0/config"
	"github.com/pkg/errors"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type DB struct {
	*bun.DB
}

func New(ctx context.Context) (*DB, error) {
	dsn := config.Get().PgDSN
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	db := bun.NewDB(sqldb, pgdialect.New())

	// db.AddQueryHook(bundebug.NewQueryHook(
	// 	bundebug.WithVerbose(true),
	// 	bundebug.FromEnv("BUNDEBUG"),
	// ))

	return &DB{db}, nil
}

func (db *DB) Reset(ctx context.Context) error {
	_, _ = db.NewDropTable().Model((*DBDelivery)(nil)).IfExists().Exec(ctx)
	_, _ = db.NewDropTable().Model((*DBPayment)(nil)).IfExists().Exec(ctx)
	_, _ = db.NewDropTable().Model((*DBItem)(nil)).IfExists().Exec(ctx)
	_, _ = db.NewDropTable().Model((*DBOrder)(nil)).IfExists().Exec(ctx)

	return nil
}

func (db *DB) Ensure(ctx context.Context) error {
	_, err := db.NewCreateTable().Model((*DBDelivery)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "pg.create")
	}

	_, err = db.NewCreateTable().Model((*DBPayment)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "pg.create")
	}

	_, err = db.NewCreateTable().Model((*DBItem)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "pg.create")
	}

	_, err = db.NewCreateTable().Model((*DBOrder)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "pg.create")
	}

	return nil
}
