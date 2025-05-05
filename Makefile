COVERAGE_FILE ?= coverage.out

.PHONY: utest
utest:
	@cd server/app && go test -coverprofile ${COVERAGE_FILE} -short ./... && go tool cover -func ${COVERAGE_FILE}

.PHONY: lint
lint:
	@cd server/app && golangci-lint run

.PHONY: itest
itest:
	@cd server/app && go test -coverprofile ${COVERAGE_FILE} -v -tags=integration ./... && go tool cover -func ${COVERAGE_FILE}

.PHONY: test-env-up
test-env-up:
	@docker compose up --exit-code-from migrate migrate

.PHONY: test-env-down
test-env-down:
	@docker compose down -v

.PHONY: fmt
fmt:
	@cd server/app && go fmt ./...
