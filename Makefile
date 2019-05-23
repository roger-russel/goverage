.PHONY: packages

packages:
	@dep ensure

.PHONY: coverage test

test:
	@go test -cover -coverprofile=./coverage.dirty.txt -covermode=atomic -coverpkg=all ./...
	@goverage clean coverage.dirty.txt -o coverage.txt --remove-origin


coverage: test
	@go tool cover -html=./coverage.txt -o coverage.html
	@google-chrome coverage.html
