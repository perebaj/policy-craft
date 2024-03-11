from typing import Union

from fastapi import FastAPI
from postgres.db import connect

app = FastAPI()

@app.get("/")
async def read_root():
    aconn = await connect()
    async with aconn.cursor() as acur:
        result = await acur.execute("SELECT 1")
        await result.fetchone()
        print(result._rowcount)
        return {"row_count": result._rowcount}
