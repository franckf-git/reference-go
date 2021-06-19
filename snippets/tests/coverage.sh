#! /bin/sh

go test -cover

# html report
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
