package controller

import (
	"context"
	"log"
	"time"

	stan "github.com/nats-io/stan.go"
	"github.com/pgeowng/wb-l0/config"
	"github.com/pgeowng/wb-l0/model"
	"github.com/pgeowng/wb-l0/service"
)

type NatsController struct {
	sc  stan.Conn
	sub stan.Subscription

	srv service.OrderService
	log *log.Logger
}

func NewNats(ctx context.Context, srv service.OrderService, log *log.Logger) (nats *NatsController, err error) {
	cfg := config.Get()

	sc, err := stan.Connect(cfg.NatsClusterId, cfg.NatsClientId, stan.ConnectWait(10*time.Second))
	if err != nil {
		return
	}

	n := &NatsController{
		sc:  sc,
		srv: srv,
		log: log,
	}

	sub, err := sc.Subscribe(cfg.NatsSubject, n.Handler)
	if err != nil {
		sc.Close()
		return
	}

	n.sub = sub
	nats = n
	return
}

func (c *NatsController) Close() {
	c.sub.Unsubscribe()
	c.sc.Close()
}

func (c *NatsController) Handler(m *stan.Msg) {
	order := model.Order{}
	err := order.FromJSONBuffer(m.Data)
	if err != nil {
		log.Printf("bad json: %s\n%s", err, string(m.Data))
		return
	}

	err = order.Validate()
	if err != nil {
		log.Printf("not valid: %s", err)
		return
	}

	log.Println(order)

	ctx := context.Background()
	err = c.srv.Create(ctx, &order)
	if err != nil {
		log.Printf("not valid: %s", err)
		return
	}
}
