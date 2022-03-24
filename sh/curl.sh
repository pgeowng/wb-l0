#!/bin/sh

PORT="1314"
API_ADDR="127.0.0.1:$PORT"

curl -s -X GET "$API_ADDR/orders"
curl -s -X GET "$API_ADDR/orders/"
