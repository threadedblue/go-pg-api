#!/bin/bash

PID_FILE=".go-api.pid"

case "$1" in
  start)
    echo "Starting Postgres..."
    docker compose up -d

    echo "Starting Go API..."
    go run ./cmd/api &
    PID=$!

    echo $PID > $PID_FILE
    echo "Go API PID: $PID"
    ;;

  stop)
    if [ -f "$PID_FILE" ]; then
        PID=$(cat $PID_FILE)
        echo "Stopping Go API (PID $PID)..."
        kill $PID
        rm $PID_FILE
    else
        echo "No PID file found."
    fi

    echo "Stopping Postgres..."
    docker compose down
    ;;

  restart)
    $0 stop
    sleep 1
    $0 start
    ;;

  *)
    echo "Usage: $0 {start|stop|restart}"
    ;;
esac
