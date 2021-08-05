#!/bin/sh

# this runs from docker-compose in the app-for-testing service

 cd /go-with-compose
# https://blog.jbowen.dev/2019/08/using-go-flags-in-tests/
go test -v ./... -endpoint=http://accountapi:8080/v1/organisation/accounts

#  go run cmd/main.go --endpoint=http://accountapi:8080/v1/organisation/accounts

# ginkgo not found in container
# ginkgo -v
