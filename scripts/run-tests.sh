#!/bin/sh

# this runs from docker-compose in the app-for-testing service

 cd /go-with-compose
# https://blog.jbowen.dev/2019/08/using-go-flags-in-tests/
# https://ieftimov.com/post/testing-in-go-go-test/
go test ./... -cover -v -endpoint=http://accountapi:8080/v1/organisation/accounts

# run this locally to view coverage
# go test ./... -cover -v -coverprofile=prof.out
# go tool cover -html=prof.out
