#!/usr/bin/env sh

# shellcheck disable=SC2164
cd "$(cd "$(dirname "$0")";pwd)"

env=$1
unit_test=$2
lint=$3
bin=car_server

set -e
export GO111MODULE=on

version=$(grep "^Version " CHANGELOG.md | head -1 | awk -F' ' '{print $2}')
fourthVersion=$(date "+%Y%m%d%H%M")
version=${version}_${fourthVersion}

if echo "${env}" | grep -q "win"; then
    export GOARCH=amd64 GOOS=windows
fi

if echo "${env}" | grep -q "linux"; then
    export GOARCH=amd64 GOOS=linux
fi

if echo "${env}" | grep -q "mac"; then
    export GOARCH=amd64 GOOS=darwin
fi

if [ "$lint" = "lint" ]; then
    golangci-lint run
fi

go_test() {
    mkdir -p coverage
    go test -coverprofile=coverage/cover.out ./...
    if [ $? -ne 0 ]; then
        echo "unit test failed!!!"
        exit 1
    fi
    go tool cover -html=coverage/cover.out -o coverage/coverage.html
}

echo ["Unit Test"]

if [ "$unit_test" = "test" ]; then
    go_test
else
    echo -e "\033[31mUnit tests are not performed!\033[0m"
fi

echo
echo [version]
go version
echo version ${version}

echo
echo [branch]
echo -n ${bin}": "
branch=$(git branch | awk '$1 == "*"{print $2}')
echo -e "\033[32m${branch}\033[0m"

echo
echo [build]
mkdir -p ./bin
go build -ldflags "-s -w -X main.version=${version}" -o ./bin/${bin} ./

echo
echo -e "\033[32m编译成功\033[0m"
echo