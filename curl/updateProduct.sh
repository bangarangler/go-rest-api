#!/bin/zsh

curl -X PUT http://localhost:8010/product/2 \
  -H "Content-Type: application/json" \
  -d "{\"name\": \"updated epic product\", \"price\": 70.00}" \
  -v | jq # need jq installed to pretty print json or remove the | jq

