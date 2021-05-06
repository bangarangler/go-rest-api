#!/bin/zsh

curl -X POST http://localhost:8010/product \
  -H "Content-Type: application/json" \
  -d "{\"name\": \"ryan\", \"price\": 1.00}" \
  -v | jq # need jq installed to pretty print json or remove the | jq

