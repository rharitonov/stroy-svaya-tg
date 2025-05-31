#!/bin/bash

NAME=$1
if [ -z "$NAME" ]; then
  echo "Использование: create_migration.sh имя_миграции"
  exit 1
fi

COUNT=$(ls migrations/*.up.sql | wc -l | awk '{print $1}')
NEXT=$(printf "%06d" $((COUNT + 1)))

echo "-- up" > "migrations/${NEXT}_${NAME}.up.sql"
echo "-- down" > "migrations/${NEXT}_${NAME}.down.sql"

echo "Созданы файлы:"
echo "- migrations/${NEXT}_${NAME}.up.sql"
echo "- migrations/${NEXT}_${NAME}.down.sql"