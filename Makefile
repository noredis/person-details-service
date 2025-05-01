COVERAGE_FILE ?= coverage.out

.PHONY: utest
utest:
	@go test -coverprofile ${COVERAGE_FILE} -short ./... && go tool cover -func ${COVERAGE_FILE}

.PHONY: lint
lint:
	@golangci-lint run

.PHONY: itest
itest:
	@go test -coverprofile ${COVERAGE_FILE} -v -tags=integration ./... && go tool cover -func ${COVERAGE_FILE}

