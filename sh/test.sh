#!/bin/sh

CLIENT="$GOPATH/src/github.com/nats-io/stan.go/examples/stan-pub/main.go"

CLUSTER="NATS"
SUBJECT="orders"

PORT="1318"
API_ADDR="localhost:$PORT"

NATS_CONTAINER=""
PG_CONTAINER=""
WB_CONTAINER=""

REQUEST=""
ORDER_UID=""
RESPONSE=""
STATUS_CODE=""

test_get_valid() {
  # $1 - index inside data.json
  # returns: REQUEST

  REQUEST="$(
    jq -c '.[]' ./sh/data.json | \
    head -n "$1" | tail -n 1 \
  )"
}

test_send() {
  # $1 nats message
  # returns: ORDER_ID

  ORDER_ID="$(echo "$REQUEST" | jq -r '.order_uid')"
  go run "$CLIENT" -c "NATS" "$SUBJECT" "$1"
}

test_exists() {
  # $1 url
  # returns: STATUS_CODE RESPONSE

  TMP="$(mktemp)"
  STATUS_CODE=$(curl -s -o "$TMP" \
    -w "%{http_code}" \
    "$1")

  RESPONSE="$(cat $TMP)"
}

test_status_code() {
  # $1 expected status code
  if [[ "$STATUS_CODE" -ne "$1" ]]; then
    echo "$TEST_MSG"
    echo "test_status_code:"
    echo "expected $1, got: $STATUS_CODE"
    echo "$RESPONSE"
    exit
  fi
}

test_entry_mismatch() {
  # $1 expected order_uid

  RES=$(curl -s -X GET "$API_ADDR/api/posts/$1")
  GOT_ID="$(echo "$RESPONSE" | jq -r '.order_uid')"

  if [ "$GOT_ID" != "$1" ]; then
    echo "$TEST_MSG"
    echo "test_entry_mismatch"
    echo "expected $1, got: $GOT_ID"
    echo "$RESPONSE" | jq
    exit
  fi
}

test_fill_db() {
  jq -c '.[]' ./sh/data.json | \
  while read line; do sleep 1; echo "$line"; done | \
  head -n 24 | \
  while IFS= read -r request; do \
    go run "$CLIENT" -c "NATS" "$SUBJECT" "$request"; \
    done
}

sudo docker container stop nats
sudo docker container rm nats

sudo docker run \
  -d \
  --name nats \
  --network pg-network \
  --network-alias nats \
  -p 4222:4222 \
  -p 8222:8222 \
  nats-streaming \
  --cluster_id NATS


sudo docker container stop pg-13
sudo docker container rm pg-13
sudo docker run -d \
  --name pg-13 \
  --network pg-network \
  --network-alias pgsql \
  -p 5432:5432 \
  -e POSTGRES_USER=sunny \
  -e POSTGRES_PASSWORD=1 \
  -e POSTGRES_DB=wbl0 \
  postgres:13.3

echo "waiting db init..."
sleep 7

sudo docker container stop wbl0
sudo docker container rm wbl0
sudo docker run \
    -d \
    --name wbl0 \
    --network pg-network \
    -p 1318:1318 \
    -e NATS_URL="nats://nats:4222" \
    -e NATS_CLUSTER_ID=NATS \
    -e NATS_CLIENT_ID=aac \
    -e NATS_SUBJECT=orders \
    -e PG_DSN="postgres://sunny:1@pg-13:5432/wbl0?sslmode=disable" \
    -e PG_RESET="false" \
    -e PORT=1318 \
    -e LOG_FILE="" \
    wb-l0

TEST_MSG="test add order"
test_get_valid 1
test_send "$REQUEST"
sleep 0.1
test_exists "$API_ADDR/orders/$ORDER_ID"
test_status_code "200"
test_entry_mismatch "$ORDER_ID"
echo "success: $TEST_MSG"

TEST_MSG="test not found"
test_exists "$API_ADDR/orders/random_string"
test_status_code "404"
echo "success: $TEST_MSG"

sudo docker container stop wbl0
sudo docker container stop pg-13
sudo docker container stop nats
