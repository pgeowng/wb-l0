#!/bin/sh

# docker pull nats-streaming
sudo docker container stop nats
sudo docker container rm nats

sudo docker run \
  -it \
  --name nats \
  --network pg-network \
  --network-alias nats \
  -p 4222:4222 \
  -p 8222:8222 \
  nats-streaming \
  --cluster_id NATS
