
# How to start to develop locally

- Install the `pipx` package manager. [Documentation](https://github.com/pypa/pipx)
- Install the `poetry` package manager. [Documentation](https://python-poetry.org/docs/)
- Using poetry to set up the project:
    - Inside the `backend` directory, run `poetry config virtualenvs.in-project true` to create a virtual environment inside the project directory.

    - Install the dependencies with `poetry install`
    - Activate the virtual environment with `poetry shell`

# Command line interface

- Run `make help` to see the available commands.
