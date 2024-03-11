from typing import Union

from fastapi import FastAPI, Depends
from contextlib import asynccontextmanager
from postgres.db import run_migration, get_db
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.sql import text

@asynccontextmanager
async def lifespan(app: FastAPI):
    yield
    try:
        run_migration()
    except Exception as e:
        print("Error running migration", e)
        pass

app = FastAPI(lifespan=lifespan)

@app.get("/")
async def read_root(db: AsyncSession = Depends(get_db)) -> Union[str, int]:
    # select 1 in the policies table
    cursor_result = await db.execute(text("SELECT 1"))
    print(cursor_result.fetchone())
    return "hello"
