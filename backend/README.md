# Getting started

To start the backend, run the following command: `make dev/start` and access the backend at `http://localhost:8080`.

to stop the backend, run the following command: `make dev/stop`

# Command line

All commands that deal with the backend are simplified in the `Makefile`. Just run `make help` to see all available commands.

# Running tests and linter

All commands could be run inside the docker container. So, to run the tests, just run `make dev/test` or the linter with `make dev/lint`.
You don't need to have any pre-requisite installed on your machine.

# Documentation

The documentation of the API can be found at api/docs. To access it, just start the backend (`make dev/start`) and access `http://localhost:8080/docs`.
