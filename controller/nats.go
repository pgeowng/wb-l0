package controller

import (
	"context"
	"fmt"
	"time"

	stan "github.com/nats-io/stan.go"
	"github.com/pgeowng/wb-l0/config"
	"github.com/pgeowng/wb-l0/service"
)

type NatsController struct {
	sc  stan.Conn
	sub stan.Subscription

	srv service.OrderService
}

func NewNats(ctx context.Context, srv service.OrderService) (nats *NatsController, err error) {
	cfg := config.Get()

	sc, err := stan.Connect(cfg.NatsClusterId, cfg.NatsClientId, stan.ConnectWait(10*time.Second))
	if err != nil {
		return
	}

	n := &NatsController{sc: sc}
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
	fmt.Printf("Received a message: %s\n", string(m.Data))
}
