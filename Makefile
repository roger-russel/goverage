.PHONY: dev

dev:
	@go run cmd/goverage/main.go -c ./tests/crud/coverage.txt


