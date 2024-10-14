#!/usr/bin/env sh

cd "$(cd "$(dirname "$0")";pwd)"

bin=car_server

echo
go version

echo
echo [build]
mkdir -p ./bin
CGO_ENABLED=0 go build -o ./bin/${bin} ./

echo
echo -e "\033[32mcompile successfully\033[0m"
echo