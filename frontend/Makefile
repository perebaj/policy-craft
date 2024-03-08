## Install all packages
.PHONY: install
install:
	npm install

## Start the development server locally
.PHONY: start
start:
	npm run dev

## Run the linter
.PHONY: lint
lint:
	npx eslint ./src/

## Build the project
.PHONY: build
build:
	npm run build

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
