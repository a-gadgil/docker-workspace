from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI(
    title="Python Enterprise Service",
    version="0.1.0",
    description="Core backend telemetry & processing service",
)


class HealthResponse(BaseModel):
    status: str
    service: str


@app.get("/health", response_model=HealthResponse)
def health_check() -> HealthResponse:
    return HealthResponse(status="healthy", service="python-app")


@app.get("/")
def root() -> dict[str, str]:
    return {"message": "Welcome to the enterprise monorepo python service"}
