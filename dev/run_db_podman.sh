#!/bin/bash

set -o nounset errexit

OUTPUT_FILE=dev/db.out
ERROR_FILE=dev/db.err

rm -f $OUTPUT_FILE $ERROR_FILE
 
export PGDATA=/var/lib/postgresql/data
nohup podman run -it --rm \
    -p 5432:5432 \
    -e POSTGRES_DB=admin \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD=pwd \
    -e "PGDATA=$PGDATA" \
    -v "./.postgres-data:$PGDATA" \
    -v ./dev/db:/docker-entrypoint-initdb.d \
    postgres \
    > "$OUTPUT_FILE" \
    2> "$ERROR_FILE" \
    &

set +o nounset errexit
