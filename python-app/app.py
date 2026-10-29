from fastapi import FastAPI
import os

app = FastAPI(title="Production Secure Service")

@app.get("/")
def read_root():
    return {
        "status": "healthy",
        "environment": os.getenv("APP_ENV", "production"),
        "user": os.getenv("USER", "appuser")
    }

if __name__ == "__main__":
    import uvicorn
    uvicorn.run("main:app", host="0.0.0.0", port=8000, reload=False)
