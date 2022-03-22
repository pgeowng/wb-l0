#!/bin/sh

CLIENT="$GOPATH/src/github.com/nats-io/stan.go/examples/stan-pub/main.go"

CLUSTER="NATS"
SUBJECT="orders"

go run "$CLIENT" \
  -c "NATS" \
 "$SUBJECT" "hello world"
