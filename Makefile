.PHONY: dev

simple:
	@go run cmd/goverage/main.go -c ./tests/goverage-test-crud/coverage.txt -p ./tests -o ./tmp/coverage.html

dev:
	@fresher -c .fresher.yaml

.PHONY: help
help:
	@go run cmd/goverage/main.go

.PHONY: print
print:
	@google-chrome-stable --headless --screenshot  --window-size=auto,auto --default-background-color=0 ./tmp/coverage.html

.PHONY: snapshot
snapshot:
	@goreleaser --snapshot --skip-publish --rm-dist
