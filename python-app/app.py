import os
import redis
import psycopg2
from flask import Flask

app = Flask(__name__)

DATABASE_URL = os.getenv("DATABASE_URL")
REDIS_URL = os.getenv("REDIS_URL")

@app.route("/")
def index():
    # Test Postgres Connection
    db_status = "Connected"
    try:
        conn = psycopg2.connect(DATABASE_URL)
        conn.close()
    except Exception as e:
        db_status = f"Error: {e}"

    # Test Redis Connection
    redis_status = "Connected"
    try:
        r = redis.Redis.from_url(REDIS_URL)
        r.ping()
    except Exception as e:
        redis_status = f"Error: {e}"

    return f"<h2>4-Tier Stack Operational!</h2><p>PostgreSQL: {db_status}</p><p>Redis: {redis_status}</p>"

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)