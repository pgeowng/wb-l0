#!/bin/sh

export NATS_CLUSTER_ID="NATS"
export NATS_CLIENT_ID="aab"
export NATS_SUBJECT="orders"

go run .