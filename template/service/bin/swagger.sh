#!/bin/bash

source "$(dirname "$0")"/utils.sh

if ! commandExist go;
then
  echo 'please install golang'
  exit 1
fi

project_dir="$(cd -- "$(dirname -- "$0")/.." &>/dev/null && pwd -P)"

cd $project_dir

export PATH=$PATH:$(go env GOPATH)/bin

if ! commandExist swag;
then
  go install github.com/swaggo/swag/cmd/swag@latest
fi

swag init -g ./internal/entrypoints/httpd/httpserver.go -d ./