from typing import Union

from fastapi import FastAPI

app = FastAPI()

@app.get("/")
def read_root() -> Union[str, int]:
    return "Hello, world! 2"
