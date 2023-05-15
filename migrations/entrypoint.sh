#!/bin/bash

set -e

DSN="host=$DB_HOST user=$POSTGRES_USER password=$POSTGRES_PASSWORD dbname=$POSTGRES_DB sslmode=disable"

goose postgres "$DSN" up
goose postgres "$DSN" status