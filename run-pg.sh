#!/bin/bash

case "$1" in
  start)
    echo "Starting PG..."
    docker compose up -d
    ;;
  stop)
    echo "Stopping PG..."
    docker compose down
    ;;
  restart)
    docker compose down
    docker compose up -d
    ;;
  logs)
    docker compose logs -f
    ;;
  *)
    echo "Usage: $0 {start|stop|restart|logs}"
    ;;
esac
