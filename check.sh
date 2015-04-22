#!/bin/sh

go build -o mer && go vet ./ && golint ./ && go test