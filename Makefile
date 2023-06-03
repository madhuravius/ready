deps:
	go mod download
.PHONY: deps

init: deps
	go generate ./...
.PHONY: init

build-image:
	docker build -f ./Containerfile -t ghcr.io/madhuravius/ready:latest .
.PHONY: build-image

build:
	go build -ldflags="-s -w" -o ./bin/ready main.go
.PHONY: build

lint:
	docker run \
		--rm \
		-v $(shell pwd):/app \
		-w /app \
		docker.io/golangci/golangci-lint:v1.52 \
		golangci-lint run
.PHONY: lint

format:
	go fmt ./...
.PHONY: format

test:
	go test ./... -v -coverprofile=coverage.out
.PHONY: test
