#!/bin/sh

# this runs from docker-compose in the app-for-testing service

# Ideas for waiting until things are up
# https://docs.docker.com/compose/startup-order/

 cd /go-with-compose
# go test  ./... 
 go run main.go
 