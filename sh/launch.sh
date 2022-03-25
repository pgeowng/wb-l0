#!/bin/sh

# export NATS_CLUSTER_ID="NATS"
# export NATS_CLIENT_ID="aag"
# export NATS_SUBJECT="orders"

# export PG_DSN="postgres://sunny:1@localhost:5432/wbl0?sslmode=disable"
# export PG_RESET="false"
# # export PG_RESET="true"

# export PORT="1318"
# export LOG_FILE="logfile"

# go run .
sudo docker build -t wb-l0 .

sudo docker container stop wbl0
sudo docker container rm wbl0
sudo docker run \
    -it \
    --name wbl0 \
    --network pg-network \
    -p "1318:1318" \
    -e NATS_URL="pg-network://nats:4222" \
    -e NATS_CLUSTER_ID=NATS \
    -e NATS_CLIENT_ID=aac \
    -e NATS_SUBJECT=orders \
    -e PG_DSN="postgres://sunny:1@pgsql:5432/wbl0?sslmode=disable" \
    -e PG_RESET="false" \
    -e PORT=1318 \
    -e LOG_FILE="logfile" \
    wb-l0
