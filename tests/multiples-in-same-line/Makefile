.PHONY: test coverage
test:
	@go test -v ./... -cover -coverprofile=./coverage.txt -covermode=atomic

coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
	@google-chrome ./coverage.html
