package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/pgeowng/wb-l0/config"
	"github.com/pgeowng/wb-l0/controller"
	"github.com/pgeowng/wb-l0/service"
	"github.com/pgeowng/wb-l0/store"
	"github.com/pkg/errors"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	go func() {
		select {
		case <-c:
			cancel()
		case <-ctx.Done():
		}
	}()

	err := launch(ctx)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}
}

func launch(ctx context.Context) error {
	st, err := store.New(ctx)
	if err != nil {
		return errors.Wrap(err, "store.init")
	}

	if config.Get().PgReset {
		err = st.DB.Reset(ctx)
		if err != nil {
			return errors.Wrap(err, "store.pg.reset")
		}
	}

	srv := service.New(ctx, st)

	nats, err := controller.NewNats(ctx, srv, log.Default())
	if err != nil {
		return errors.Wrap(err, "controller.nats")
	}

	defer nats.Close()

	<-ctx.Done()

	return nil
}
