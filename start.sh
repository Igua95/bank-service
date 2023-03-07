#!/bin/sh

# Exit
set -e

echo "Run DB migration"
source /app/app.env
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Start app"
# means take all parameters an run it
exec "$@"