#!/bin/sh

# init
go mod init app
go mod tidy

# run
go run *.go

# test
go test ./animals
go test .