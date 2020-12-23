#!/bin/bash

go build main.go
mv main main_darwin_amd64
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
mv main main_linux_amd64
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
mv main.exe main_windows_amd64.exe
go build main.go
