# Getting Started

- Run `make dev/start` to start the backend. and Access the API at `http://localhost:8000/` or the API documentation at `http://localhost:8000/docs`

Obs: `make dev/stop` to stop the backend.

# How to run the tests and linters

Run `make dev/test` to run the tests. Or `make dev/lint` to run the linter. Simple as that.

# How to start to develop locally

To start developing locally, you need to have the following tools installed:

- Install the `pipx` package manager. [Documentation](https://github.com/pypa/pipx)
- Install the `poetry` package manager. [Documentation](https://python-poetry.org/docs/)
- Using poetry to set up the project:
    - Inside the `backend` directory, run `poetry config virtualenvs.in-project true` to create a virtual environment inside the project directory.

    - Install the dependencies with `poetry install`
    - Activate the virtual environment with `poetry shell`

# Command line interface

Some useful commands are available to help you with the development process:

- Run `make help` to see the available commands.
