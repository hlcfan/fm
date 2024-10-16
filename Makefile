.PHONY: build

build:
	go build -ldflags="-s -w" -o fm main.go

test:
	go test ./... -v
