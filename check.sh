#!/bin/sh

go build -o mer
golint ./
go vet ./
go test