import psycopg
import os

postgres_url = os.environ.get("POLICYCRAFT_POSTGRES_URL")
if postgres_url is None:
    raise ValueError("POLICYCRAFT_POSTGRES_URL must be set")

"""
connect return a async connection to the postgres database
@return a async connection to the postgres database
"""
async def connect() -> psycopg.AsyncConnection:
    aconn = await psycopg.AsyncConnection.connect(
        conninfo=postgres_url,
        autocommit=False
    )
    return aconn
