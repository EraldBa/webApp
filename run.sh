#!/usr/bin/fish

go build -o fitbuddy -ldflags="-s -w" cmd/web/*.go && ./fitbuddy