#!/usr/bin/env bash
#This file helps to know code-coverage so as to trigger fast development.
go test ./... -v -coverprofile=coverage.out.tmp -p 1
cat coverage.out.tmp | grep -v "_mock.go" >coverage.out
go tool cover -func=coverage.out
