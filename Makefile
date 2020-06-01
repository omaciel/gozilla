all: fmt lint vet build

fmt:
	@echo "Formatting the source code"
	go fmt ./...

lint:
	@echo "Linting the source code"
	golint ./...

vet:
	@echo "Checking for code issues"
	go vet ./...

clean:
	@echo "Removing binaries"
	rm -rf bin

build: clean
	@echo "Building the client binary"
	go build -o bin/gozilla *.go
