IMG ?= docker.io/astropulseinc/latency-app
TAG ?= dev

.PHONY: build run test docker
build: ## Build the binary into bin/
	go build -o bin/latency .

run: build ## Run on http://localhost:8080
	./bin/latency

test: ## Vet and test
	go vet ./... && go test ./...

docker: ## Build the image for this machine as $(IMG):$(TAG)
	docker build -t $(IMG):$(TAG) .
