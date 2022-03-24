#!/bin/sh

# docker run --name nats \
#   --network nats \
#   --rm \
#   -p 4222:4222 \
#    -p 8222:8222 \
#   nats \
#   --http_port 8222 \
#   --cluster_name NATS \
#   --cluster nats://0.0.0.0:6222

# docker pull nats-streaming

sudo docker run \
  -p 4222:4222 \
  -p 8222:8222 \
  nats-streaming \
  --cluster_id NATS
