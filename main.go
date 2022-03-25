package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"
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

	err = st.DB.Ensure(ctx)
	if err != nil {
		return errors.Wrap(err, "store.pg.ensure")
	}

	srv := service.New(ctx, st)

	err = srv.Recover(ctx)
	if err != nil {
		return errors.Wrap(err, "srv")
	}

	llog := log.Default()

	if len(cfg.LogFile) > 0 {
		f, err := os.OpenFile(cfg.LogFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		llog.SetOutput(f)
	}

	nats, err := controller.NewNats(ctx, srv, llog)
	if err != nil {
		return errors.Wrap(err, "controller.nats")
	}

	defer nats.Close()

	rest := controller.NewRest(ctx, srv, log.Default())

	engine := html.New("./views", ".html")
	engine.AddFunc("offsetInt", func(idx int, offset int) int {
		return idx + offset
	})
	engine.AddFunc("Iterate", func(count int) []int {
		var i int
		var st []int
		for i = 0; i < count; i++ {
			st = append(st, i)
		}
		return st
	})
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Use(logger.New())
	app.Get("/orders", rest.GetIds)
	app.Get("/orders/:id", rest.GetOrder)
	app.Get("/:idx?", rest.IndexPage)
	fmt.Printf("Listen on :%s\n", cfg.HttpPort)
	go func() {
		log.Println(app.Listen(":" + cfg.HttpPort))
	}()
	defer app.Shutdown()

	<-ctx.Done()

	return nil
}
