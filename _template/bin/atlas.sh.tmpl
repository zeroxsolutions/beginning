#!/bin/bash

source "$(dirname "$0")"/utils.sh

export PATH=$PATH:$(go env GOPATH)/bin

if ! commandExist atlas;
then
    curl -sSf https://atlasgo.sh | sh
fi

if [ -f .env ];
then
    export $(cat .env | xargs)
fi

atlas migrate diff --env gorm