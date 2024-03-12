from typing import Union

from fastapi import FastAPI, Depends
from postgres.db import db_session
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.sql import text

app = FastAPI()

@app.get("/")
async def read_root(session: AsyncSession = Depends(db_session)) -> Union[str, int]:
    # select 1 in the policies table
    cursor_result = await session.execute(text("SELECT 1"))
    print(cursor_result.fetchone())
    return "hello"
