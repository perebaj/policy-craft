import asyncio
import pytest
from alembic import command, config
from pytest_async_sqlalchemy import create_database, drop_database

@pytest.fixture(scope="session")
def event_loop():
    loop = asyncio.get_event_loop()
    yield loop
    loop.close()

@pytest.fixture(scope="session")
def _database_url():
    return "postgresql+asyncpg://postgres:postgres@localhost:5432/postgres"

@pytest.fixture(scope="session")
async def database(database_url, event_loop):
    await create_database(database_url, event_loop)

    alembic_cfg = config.Config("../alembic.ini")
    command.upgrade(alembic_cfg, "head")

    try:
        yield database_url
    finally:
        await drop_database(database_url, event_loop)
