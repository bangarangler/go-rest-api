#!/bin/zsh

curl -X DELETE http://localhost:8010/product/4 -v | jq # need jq installed to pretty print json or remove the | jq

