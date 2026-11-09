IMG ?= docker.io/astropulseinc/latency-app
TAG ?= dev

.PHONY: build run test docker
build:
	go build -o bin/latency .

run: build
	./bin/latency

test:
	go vet ./... && go test ./...

docker:
	docker build -t $(IMG):$(TAG) .
