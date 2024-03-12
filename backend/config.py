# config.py gather all the necessary settings for the application.
from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    # The URL to connect to the PostgreSQL database.
    POLICYCRAFT_POSTGRES_URL: str

settings = Settings()
