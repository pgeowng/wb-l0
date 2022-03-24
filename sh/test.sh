#!/bin/sh

CLIENT="$GOPATH/src/github.com/nats-io/stan.go/examples/stan-pub/main.go"

CLUSTER="NATS"
SUBJECT="orders"

# go run "$CLIENT" \
#   -c "NATS" \
#  "$SUBJECT" "hello world"

VALID="$(cat ../model/model.json | tr '\n' ' ')"

# go run "$CLIENT" -c "NATS" "$SUBJECT" '{"hello":"world"}'

jq -c '.[]' all.json | \
  while read line; do sleep 1; echo "$line"; done | \
  while IFS= read -r request; do \
    go run "$CLIENT" -c "NATS" "$SUBJECT" "$request"; \
    done
