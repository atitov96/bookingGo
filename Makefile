.PHONY: run test lint build clean

BINARY_NAME=hotel-booking-service
GO=go

run:
	@echo "Starting service..."
	$(GO) run ./cmd/api/main.go

build:
	@echo "Building..."
	$(GO) build -o $(BINARY_NAME) ./cmd/api/main.go

test:
	@echo "Running tests..."
	$(GO) test -v ./...

lint:
	@echo "Running linter..."
	$(GO) vet ./...
	@if command -v golangci-lint >/dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint is not installed"; \
		exit 1; \
	fi

clean:
	@echo "Cleaning..."
	$(GO) clean
	rm -f $(BINARY_NAME)

dev: lint test run
