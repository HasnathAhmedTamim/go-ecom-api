# Simple Makefile for common dev tasks
.PHONY: run lint build smoke fmt vet tidy

run:
	@echo "Running server..."
	JWT_SECRET=devsecret go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

fmt:
	go fmt ./...

vet:
	go vet ./...

tidy:
	go mod tidy

smoke:
	@echo "Run smoke tests against http://127.0.0.1:8080"
	go run ../tools/smoke.go
