.PHONY: test coverage
test:
	@go test -v ./... -coverpkg="./internal/...,./pkg/...,./cmd/..." -cover -coverprofile=./coverage.txt -covermode=atomic

coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
	@google-chrome ./coverage.html
