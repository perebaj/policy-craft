from fastapi import FastAPI, Depends
from postgres.db import connect
from psycopg import AsyncConnection
from contextlib import asynccontextmanager

"""
get_db is a dependency that will return a connection to the postgres database
"""
async def get_conn():
    aconn = await connect()
    try:
        yield aconn
    finally:
        await aconn.close()

@asynccontextmanager
async def lifespan(app: FastAPI):
   # run the alembic migrations and trigger an error if they fail
    print("Running migrations")
    yield

app = FastAPI(lifespan=lifespan)

@app.get("/")
async def read_root(aconn: AsyncConnection = Depends(get_conn)):
    async with aconn.cursor() as cur:
        await cur.execute("SELECT 1")
        result = await cur.fetchone()
        return {"message": f"Hello, world! {result}"}
