# db.py is responsible for creating the database session and engine
from typing import AsyncGenerator

from config import settings
from sqlalchemy.ext.asyncio import create_async_engine, AsyncSession
from sqlalchemy.orm import sessionmaker

if not settings.POLICYCRAFT_POSTGRES_URL:
    raise ValueError("POLICYCRAFT_POSTGRES_URL is not set")

async_engine = create_async_engine(settings.POLICYCRAFT_POSTGRES_URL, echo=True, future=True)

# db_session is a context manager that provides a database session to the caller.
async def db_session() -> AsyncGenerator:
    async_session = sessionmaker(
        bind=async_engine,
        class_=AsyncSession,
        expire_on_commit=False,
    )
    async with async_session() as session:
        yield session
