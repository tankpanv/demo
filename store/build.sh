#!/usr/bin/env bash
RUN_NAME="wealth-stock-mainstore"
workspace=`pwd`
mkdir -p output/bin
cp script/* output/
chmod +x output/bootstrap.sh
cd testApi
go build 
cp testApi ../output
cd $workspace
if [ "$IS_SYSTEM_TEST_ENV" != "1" ]; then
    go build -o output/bin/${RUN_NAME}
else
    go test -c -covermode=set -o output/bin/${RUN_NAME} -coverpkg=./...
fi
