#!/bin/sh

export NATS_CLUSTER_ID="NATS"
export NATS_CLIENT_ID="aab"
export NATS_SUBJECT="orders"

export PG_DSN="postgres://sunny:1@localhost:5432/wbl0?sslmode=disable"
export PG_RESET="true"

go run ..
