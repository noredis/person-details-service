COVERAGE_FILE ?= coverage.out

.PHONY: test
test:
	@go test -coverprofile ${COVERAGE_FILE} ./... && go tool cover -func ${COVERAGE_FILE}

.PHONY: lint
lint:
	@golangci-lint run

