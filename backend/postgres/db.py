
from alembic import command
from alembic.config import Config
import os
from sqlalchemy.ext.asyncio import create_async_engine, async_sessionmaker


postgres_url = os.getenv("POLICYCRAFT_POSTGRES_URL")
if postgres_url is None:
    raise ValueError("POLICYCRAFT_POSTGRES_URL is not set")


def run_migration() -> None:
    alembic_cfg = Config("alembic.ini")
    alembic_cfg.set_main_option("sqlalchemy.url", postgres_url)
    command.upgrade(alembic_cfg, "head", sql=False)

# get_db is a function that returns a database session
async def get_db():
    # engine it's related to define which database we are working with, in this case postgresql
    engine = create_async_engine(postgres_url)

    SessionLocal = async_sessionmaker(autocommit=False, autoflush=False, bind=engine)
    db = SessionLocal()
    try:
        yield db
    finally:
        await db.close()
