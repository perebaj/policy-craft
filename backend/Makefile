# TOOLS VERSIONS
export GO_VERSION=1.22.1
export GOLANGCI_LINT_VERSION=v1.56.2

devimage=policycraft-dev
gopkg=$(devimage)-gopkg
gocache=$(devimage)-gocache
devrun=docker-compose run --rm policycraft
image=perebaj

# to enable the tracking of the service version, we can use the git commit hash
version=$(shell git rev-parse --short HEAD)

## run all tests. Usage `make test` or `make test testcase="TestFunctionName"` to run an isolated tests
.PHONY: test
test:
	if [ -n "$(testcase)" ]; then \
		go test ./... -timeout 10s -race -run="^$(testcase)$$" -v; \
	else \
		go test ./... -timeout 10s -race; \
	fi

## Run all tests including the integration tests (requires docker up and running). Usage `make integration-test` or `make integration-test testcase="TestFunctionName"` to run an isolated tests
.PHONY: integration-test
integration-test:
	if [ -n "$(testcase)" ]; then \
		go test ./... -timeout 5s -tags integration -v -run="^$(testcase)$$" ; \
	else \
		go test ./... -timeout 5s -tags integration; \
	fi

## Calculate code coverage
.PHONY: coverage
coverage:
	go test ./... -coverprofile=c.out
	go tool cover -html=c.out

## builds the service
.PHONY: service
service:
	go build -o ./cmd/policycraft/policycraft ./cmd/policycraft

## runs the service locally
.PHONY: run
run: service
	./cmd/policycraft/policycraft

## lint the whole project
.PHONY: lint
lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION) run ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...


## Build the service image
.PHONY: image
image:
	docker build . \
		--build-arg GO_VERSION=$(GO_VERSION) \
		-t $(image)

## Publish the service image
.PHONY: image/publish
image/publish: image
	docker push $(image)

.PHONY: dev
dev: dev/image
	$(devrun)

## Create the dev container image
.PHONY: dev/image
dev/image:
	docker build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GOLANGCI_LINT_VERSION=$(GOLANGCI_LINT_VERSION) \
		-t $(devimage) \
		-f Dockerfile.dev \
		.

##run a make target inside the dev container
dev/%: dev/image
	$(devrun) make ${*}

## Start containers, additionaly you can provide rebuild=true to force rebuild
.PHONY: dev/start
dev/start:
	@echo "Starting development server..."
	@if [ "$(rebuild)" = "true" ]; then \
		docker-compose up -d --build; \
	else \
		docker-compose up -d; \
	fi

## Dev container logs
.PHONY: dev/logs
dev/logs:
	docker-compose logs -f policycraft

## Dev container stop
.PHONY: dev/stop
dev/stop:
	docker-compose stop

## Dev container cleanup
.PHONY: dev/cleanup
dev/cleanup:
	docker-compose down -v --remove-orphans --rmi all

## Access the container
dev:
	@$(devrun) bash

## Create a new migration, use make migration/new name=<migration_name>
.PHONY: migration/new
migration/new: param-name
	@echo "Creating new migration..."
	go run github.com/golang-migrate/migrate/v4/cmd/migrate \
		create \
		-dir ./postgres/migrations \
		-ext sql \
		-seq \
		$(name)

## Run migrations
.PHONY: migration/up
migration/up: param-POLICYCRAFT_POSTGRES_URL
	@echo "Running migration up..."
	go run --tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate \
		-path=./postgres/migrations \
		-database=$(POLICYCRAFT_POSTGRES_URL) \
		-verbose \
		up

## Step down a single migration
.PHONY: migration/down
migration/down: param-POLICYCRAFT_POSTGRES_URL
	@echo "Running migration down..."
	go run --tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate \
		-path=./postgres/migrations \
		-database=$(POLICYCRAFT_POSTGRES_URL) \
		-verbose \
		down 1

## Drop the database
.PHONY: migration/drop
migration/drop: param-POLICYCRAFT_POSTGRES_URL
	@echo "Running migration drop..."
	go run --tags postgres github.com/golang-migrate/migrate/v4/cmd/migrate \
		-path=./postgres/migrations \
		-database=$(POLICYCRAFT_POSTGRES_URL) \
		-verbose \
		drop

## Display help for all targets
.PHONY: help
help:
	@awk '/^.PHONY: / { \
		msg = match(lastLine, /^## /); \
			if (msg) { \
				cmd = substr($$0, 9, 100); \
				msg = substr(lastLine, 4, 1000); \
				printf "  ${GREEN}%-30s${RESET} %s\n", cmd, msg; \
			} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

### param-% it's an easy way to pass parameters to make targets(commands)
param-%:
	@ if [ "${${*}}" = "" ]; then \
		echo "Param \"$*\" is missing, use: make $(MAKECMDGOALS) $*=<value>"; \
		exit 1; \
	fi
