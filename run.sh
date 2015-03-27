#!/bin/bash  
echo "Building..."

pid=`lsof -P | grep ":30022" | awk '{print $2}'`

kill -9 $pid
echo "Killed port $pid"

go build
go run main.go