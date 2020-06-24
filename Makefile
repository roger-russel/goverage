.PHONY: dev

simple:
	@go run cmd/goverage/main.go -c ./tests/crud/coverage.txt -p ./tests/crud -o coverage2.html

dev:
	@fresher -c .fresher.yaml

.PHONY: help
help:
	@go run cmd/goverage/main.go
