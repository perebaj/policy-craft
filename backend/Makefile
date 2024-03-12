export PYTHON_VERSION=3.10.12
export POETRY_VERSION=1.8.2

base_image=registry.heroku.com/policy-craft/web
# version based on git commit hash to turn easy to track deployed versions
version:=$(shell git rev-parse --short HEAD)
image=$(base_image):$(version)

# command to run the development container(local environment for developers)
devimage=policy-craft
devrun=docker run -it --rm -v $(PWD):/app $(devimage)

## Run test | will run all files with test_*.py or *_test.py
.PHONY: test
test:
	poetry run pytest

## Build the service image
.PHONY: image
image:
	docker build \
		--build-arg PYTHON_VERSION=$(PYTHON_VERSION) \
		--build-arg POETRY_VERSION=$(POETRY_VERSION) \
		-t $(image) \
		.

## Publish the service image
.PHONY: image/publish
image/publish: image
	docker push $(image)

# command to build the development container, used in other commands, such as dev, dev/start, dev/stop, etc.
.PHONY: dev/image
dev/image:
	docker build \
		--build-arg PYTHON_VERSION=$(PYTHON_VERSION) \
		--build-arg POETRY_VERSION=$(POETRY_VERSION) \
		-t $(devimage) \
		-f Dockerfile.dev \
		.

## Restart containers, if container=<name> is provided only it will be restarted
.PHONY: dev/restart
dev/restart: container=
dev/restart:
	@echo "Restarting development server..."
	@docker-compose restart $(container)

## Show logs, if container=<name> is provided logs for only that container will be shown
.PHONY: dev/logs
dev/logs:
	@echo "Showing logs..."
	@docker-compose logs -f $(container)

## Start containers, additionaly you can provide rebuild=true to force rebuild
.PHONY: dev/start
dev/start:
	@echo "Starting development server..."
	@if [ "$(rebuild)" = "true" ]; then \
		docker-compose up -d --build; \
	else \
		docker-compose up -d; \
	fi

## Stop containers
.PHONY: dev/stop
dev/stop:
	@echo "Stopping development server..."
	@docker-compose down

## Clean containers, images and volumes
.PHONY: dev/clean
dev/clean:
	@echo "Cleaning containers, images and volumes..."
	@docker-compose down --rmi all --volumes --remove-orphans

## Print the IP of the running container to access the service
.PHONY: ip
ip:
	$(eval IP := $(shell docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' backend_policycraft_1))
	@echo "http://$(IP):8000"

## Access the service running container
.PHONY: dev
dev: dev/image
	$(devrun) bash

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
