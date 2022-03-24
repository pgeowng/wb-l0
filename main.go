package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	cfg := config.Get()
	st, err := store.New(ctx)
	if err != nil {
		return errors.Wrap(err, "store.init")
	}

	if cfg.PgReset {
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

	rest := controller.NewRest(ctx, srv, log.Default())

	app := fiber.New()
	app.Use(logger.New())
	app.Get("/orders", rest.GetIds)
	app.Get("/orders/:id", rest.GetOrder)
	fmt.Printf("Listen on :%s\n", cfg.HttpPort)
	go app.Listen(":" + cfg.HttpPort)
	defer app.Shutdown()

	<-ctx.Done()

	return nil
}
