#!/bin/sh

openapi-generator generate -i specification/beacon.yaml -g go-gin-server
