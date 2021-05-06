#!/bin/zsh

curl -X POST http://localhost:8010/product \
  -H "Content-Type: application/json" \
  -d "{\"name\": \"to be deleted\", \"price\": 8.00}" \
  -v | jq # need jq installed to pretty print json or remove the | jq

