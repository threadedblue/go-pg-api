#!/usr/bin/env bash

docker restart pg
docker exec -it pg psql -U postgres -d postgres -c "select current_user, current_database();"
