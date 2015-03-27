#!/bin/bash  
echo "Building..."

pid=`lsof -P | grep ":3000" | awk '{print $2}'`

kill -9 $pid
echo "Killed port $pid"

go build
go run main.go