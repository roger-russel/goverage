.PHONY: dev

simple:
	@go run cmd/goverage/main.go -c ./tests/goverage-test-crud/coverage.txt -p ./tests -o ./tmp/coverage.html

dev:
	@fresher -c .fresher.yaml

.PHONY: help
help:
	@go run cmd/goverage/main.go
