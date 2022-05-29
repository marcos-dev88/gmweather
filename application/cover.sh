#!/bin/bash

go test -v -failfast -count=1 -coverprofile=coverage.out;
go tool cover -html=coverage.out
