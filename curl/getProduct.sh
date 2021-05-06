#!/bin/zsh

curl http://localhost:8010/product/2 -v | jq # need jq installed to pretty print json or remove the | jq

