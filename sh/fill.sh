#!/bin/sh

CLIENT="$GOPATH/src/github.com/nats-io/stan.go/examples/stan-pub/main.go"
CLUSTER="NATS"
SUBJECT="orders"

jq -c '.[]' ./sh/data.json | \
  while read line; do sleep 0.24; echo "$line"; done | \
  while IFS= read -r request; do \
    go run "$CLIENT" -c "NATS" "$SUBJECT" "$request"; \
    done

#   head -n 24 | \
