#!/bin/sh

# init
go mod init chapter2
go mod tidy

# run
go run *.go

# test
go test .
go test ./animals