include local.env
export

DATABASE_CONTAINER=soat1_challenge1
DATABASE_IMAGE=${DATABASE_CONTAINER}:${DATABASE_TAG}
MOD_NAME := $(shell head -1 go.mod | sed -E 's/module //')
BENCHMARK_PACKAGES := $(shell grep 'func Bench' . -r --include \*.go | sed -E 's@\./(.*)/[a-z_]+\.go:.*@${MOD_NAME}/\1@' | sort | uniq)

.DEFAULT_TARGET: help

.PHONY: help
help: ## Display this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: run
run: ## Run HTTP server locally on port 8080
	@go run cmd/main.go

.PHONY: lint
lint: ## Execute syntatic analysis in the code and autofix minor problems
	@golangci-lint run --fix

.PHONY: test
test: ## Execute the tests in the development environment
	@go test ./... -count=1 -race -timeout 2m

.PHONY: coverage
coverage: ## Generate test coverage in the development environment
	@go test ./... -coverprofile=/tmp/coverage.out -coverpkg=./...
	@go tool cover -html=/tmp/coverage.out

.PHONY: db-build
db-build: ## Builds a mysql docker image locally
	docker build -f Dockerfile.mysql -t ${DATABASE_IMAGE} .

.PHONY: db-stop
db-stop: ## Stop database container
	docker stop ${DATABASE_CONTAINER}

.PHONY: db-run
db-run: ## Runs the mysql docker container on localhost:3306
	-docker rm -f ${DATABASE_CONTAINER}
	docker run -d -p 3306:3306 -p 33060:33060 --name ${DATABASE_CONTAINER} ${DATABASE_IMAGE}

.PHONY: db-reset
db-reset: ## Reset table registers to initial state (with seeds)
	docker exec -d ${DATABASE_CONTAINER} mysql -u admin -p${DATABASE_PASSWORD} ${DATABASE_NAME} -e 'source /docker-entrypoint-initdb.d/it_tests_seed.sql'

.PHONY: db-logs
db-logs: ## Show database logs
	docker logs -f --tail 100 ${DATABASE_CONTAINER}

.PHONY: db-bash
db-bash: ## Open bash terminal on database container
	docker exec -it ${DATABASE_CONTAINER} bash
