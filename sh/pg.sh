#!/bin/sh

sudo docker run -it \
  --rm \
  --name pg-13 \
  --network pg-network \
  -p 5432:5432 \
  -e POSTGRES_USER=sunny \
  -e POSTGRES_PASSWORD=1 \
  -e POSTGRES_DB=wbl0 \
  postgres:13.3
