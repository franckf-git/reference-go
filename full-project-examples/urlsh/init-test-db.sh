#!/bin/sh

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE test_$POSTGRES_DB;
    GRANT ALL PRIVILEGES ON DATABASE test_$POSTGRES_DB TO $POSTGRES_USER;
EOSQL
