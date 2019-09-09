#!/bin/sh
set -e

openapi-generator generate -i specification/beacon.yaml -g go-gin-server
cat <<EOF
Now you need to edit the files in go/ to get them working again
EOF
