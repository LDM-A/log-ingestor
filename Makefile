test: build
	@go test ./...

build:
	@go build -o bin/log-ingestor

run: build
	@./bin/log-ingestor