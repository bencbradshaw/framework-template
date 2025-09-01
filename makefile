.PHONY: run build dev install clean

# Development server with auto-reload
dev:
	go run main.go --dev

# Production server
run:
	go run main.go

# Build frontend and backend for production
build:
	go run main.go build

# Install dependencies
install:
	go mod download
	cd frontend && npm install

# Clean build artifacts
clean:
	rm -rf static/
	rm -f main

# Production build binary
build-binary:
	go build -o main main.go