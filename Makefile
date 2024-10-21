BUILD_DIR=build
BINARY_NAME=dodl
PKG=./...
PLATFORMS=linux darwin windows
ARCHITECTURES=amd64 arm64

default: build

build: test
	@echo "Building for all platforms and architectures..."
	@for GOOS in $(PLATFORMS); do \
		for GOARCH in $(ARCHITECTURES); do \
			echo "Building for $$GOOS/$$GOARCH..."; \
			mkdir -p $(BUILD_DIR) && \
			GOOS=$$GOOS GOARCH=$$GOARCH go build -o $(BUILD_DIR)/$(BINARY_NAME)-$$GOOS-$$GOARCH; \
		done; \
	done

test: vet
	@echo "Running tests..."
	go test $(PKG) -v

fmt:
	@echo "Running code formatting..."
	go fmt $(PKG)

lint:
	@echo "Running linter..."
	golangci-lint run

vet: fmt lint
	@echo "Running go vet..."
	go vet $(PKG)

clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

deps:
	@echo "Installing dependencies..."
	go mod tidy

install: build
	@echo "Installing the binary..."
	go install ./

.PHONY: build test fmt lint vet clean deps install
