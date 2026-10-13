import os
import redis
import psycopg2
from flask import Flask

app = Flask(__name__)

DB_HOST = os.getenv("DB_HOST", "db")
DB_NAME = os.getenv("DB_NAME", "appdb")
DB_USER = os.getenv("DB_USER", "postgres")
DB_PASSWORD = os.getenv("DB_PASSWORD", "postgres")
REDIS_HOST = os.getenv("REDIS_HOST", "redis")

def get_db_connection():
    return psycopg2.connect(
        host=DB_HOST,
        database=DB_NAME,
        user=DB_USER,
        password=DB_PASSWORD,
        connect_timeout=3
    )

def get_redis_client():
    return redis.Redis(host=REDIS_HOST, port=6379, decode_responses=True)

@app.route("/")
def index():
    try:
        conn = get_db_connection()
        cur = conn.cursor()
        cur.execute("SELECT NOW();")
        db_time = cur.fetchone()[0]
        cur.close()
        conn.close()
    except Exception as e:
        db_time = f"Database Error: {e}"

    try:
        r = get_redis_client()
        count = r.incr("hits")
    except Exception as e:
        count = f"Redis Error: {e}"

    return f"""
    🚀 Flask App Stack Running via Docker Compose!<br>
    📊 Redis Visit Count: <b>{count}</b><br>
    🗄️ PostgreSQL Database Time: <b>{db_time}</b>
    """

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080)