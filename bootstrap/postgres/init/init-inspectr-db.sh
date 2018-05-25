#!/bin/bash
set -e

DBNAME="inspectr"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	CREATE USER ${DBNAME};
  CREATE DATABASE ${DBNAME};
	GRANT ALL PRIVILEGES ON DATABASE ${DBNAME} TO ${DBNAME};
  \c ${DBNAME}
  CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
  CREATE EXTENSION IF NOT EXISTS hstore;
EOSQL