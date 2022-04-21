#!/bin/sh

set -e

wait-for "${DATABASE_HOST}:${DATABASE_PORT}" -- "$@"

/go/bin/watcher -run github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/cmd/main
