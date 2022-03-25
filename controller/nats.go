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

	urlOpt := stan.NatsURL("nats://localhost:4222")
	if len(cfg.NatsURL) > 0 {
		urlOpt = stan.NatsURL(cfg.NatsURL)
	}

	sc, err := stan.Connect(cfg.NatsClusterId, cfg.NatsClientId, stan.ConnectWait(10*time.Second), urlOpt)
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

func (ctl *NatsController) Close() {
	ctl.sub.Unsubscribe()
	ctl.sc.Close()
}

func (ctl *NatsController) Handler(m *stan.Msg) {
	var order model.Order
	msg := m.Data
	err := order.FromJSONBuffer(msg)
	if err != nil {
		ctl.log.Printf("nats.parse: %s", err)
		ctl.log.Printf("msg: %s\n", string(m.Data))
		return
	}

	err = order.Validate()
	if err != nil {
		ctl.log.Printf("nats.validate: %s\n", err)
		ctl.log.Printf("msg: %s\n", string(m.Data))
		return
	}

	ctx := context.Background()
	err = ctl.srv.Create(ctx, &order)
	if err != nil {
		ctl.log.Printf("nats.create: %s\n", err)
		ctl.log.Printf("msg: %s\n", string(msg))
		return
	}
}
