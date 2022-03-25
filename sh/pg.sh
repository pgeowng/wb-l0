#!/bin/sh


sudo docker container stop pg-13
sudo docker container rm pg-13
sudo docker run \
  -it \
  --name pg-13 \
  --network pg-network \
  --network-alias pgsql\
  -p 5432:5432 \
  -e POSTGRES_USER=sunny \
  -e POSTGRES_PASSWORD=1 \
  -e POSTGRES_DB=wbl0 \
  postgres:13.3
  # --rm \
