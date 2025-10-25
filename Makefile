TAG := $(shell git describe --tags --abbrev=0)
IMG ?= astropulseinc/latency
CONTAINER_TOOL ?= docker

UNAME_S := $(shell uname -s)
ifeq ($(UNAME_S),Linux)
    SED_INPLACE := -i
endif
ifeq ($(UNAME_S),Darwin)
    SED_INPLACE := -i ''
endif

.PHONY: docker-build
docker-build: ## Build and push docker images for multiple architectures.
	@if ! docker buildx inspect mybuilder > /dev/null 2>&1; then \
		echo "Creating new builder instance"; \
		docker buildx create --use --name mybuilder --platform linux/amd64,linux/arm64; \
	else \
		echo "Existing builder instance found, using it"; \
		docker buildx use mybuilder; \
	fi
	cp docker-bake.template.hcl docker-bake.hcl
	sed $(SED_INPLACE) 's/{{TAG}}/$(TAG)/g' docker-bake.hcl
	@echo "Building multi-platform images (this will push to registry)..."
	$(CONTAINER_TOOL) buildx bake -f docker-bake.hcl --push

.PHONY: docker-build-local
docker-build-local: ## Build docker image for current platform only (local testing).
	@echo "Building for current platform only..."
	$(CONTAINER_TOOL) build -t astropulseinc/latency-app:$(TAG) -t astropulseinc/latency-app:latest .
	@echo "✅ Image built and loaded locally as astropulseinc/latency-app:$(TAG)"

.PHONY: build
build: ## Build Go binary
	go build -o bin/latency main.go

.PHONY: run
run: build ## Run the application locally
	./bin/latency

.PHONY: docker-clean
docker-clean: ## Remove buildx builder
	docker buildx rm mybuilder

