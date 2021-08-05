#!/bin/sh

# this runs from docker-compose in the app-for-testing service

 cd /go-with-compose
go test  ./... 
#  go run cmd/main.go --endpoint=http://accountapi:8080/v1/organisation/accounts
