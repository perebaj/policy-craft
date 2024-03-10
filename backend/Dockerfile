# Dockerfile for the production service environment.
ARG PYTHON_VERSION
FROM python:${PYTHON_VERSION}-buster

ARG POETRY_VERSION

RUN pip install poetry==${POETRY_VERSION}

WORKDIR /app

COPY pyproject.toml poetry.lock ./

# as default, poetry creates a virtual environment for the project, as we are using docker, we don't need it
RUN poetry config virtualenvs.create false
# --no-dev> don't install dev dependencies
# --no-interaction> don't ask for user input
RUN poetry install --no-dev --no-interaction

COPY . /app

EXPOSE 8000

# starting the web server
CMD ["poetry", "run", "uvicorn", "app.main:app", "--host", "0.0.0.0", "--port", "8000"]
