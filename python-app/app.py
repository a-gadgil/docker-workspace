import os
import redis
import psycopg2
import requests
from flask import Flask, jsonify

app = Flask(__name__)

# Environment variables injected by Docker Compose
PG_HOST = os.getenv("PG_HOST", "localhost")
COUCH_HOST = os.getenv("COUCH_HOST", "localhost")
REDIS_HOST = os.getenv("REDIS_HOST", "localhost")

@app.route("/")
def health_check():
    status = {}

    # 1. Test Postgres Connection
    try:
        conn = psycopg2.connect(
            host=PG_HOST, database="relational_db", user="user", password="password", port=5432
        )
        cursor = conn.cursor()
        cursor.execute("SELECT version();")
        status["postgres"] = "Connected: " + cursor.fetchone()[0][:30]
        cursor.close()
        conn.close()
    except Exception as e:
        status["postgres"] = f"Error: {str(e)}"

    # 2. Test CouchDB Connection
    try:
        response = requests.get(f"http://admin:password@{COUCH_HOST}:5984/")
        if response.status_code == 200:
            status["couchdb"] = "Connected: " + response.json().get("couchDB", "OK")
        else:
            status["couchdb"] = f"Error: Status code {response.status_code}"
    except Exception as e:
        status["couchdb"] = f"Error: {str(e)}"

    # 3. Test Redis Connection & Increment Hit Counter
    try:
        r = redis.Redis(host=REDIS_HOST, port=6379, decode_responses=True)
        hits = r.incr("hits")
        status["redis"] = f"Connected: incremented hit counter to {hits}"
    except Exception as e:
        status["redis"] = f"Error: {str(e)}"

    return jsonify({"status": "success", "tier_connections": status})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8000)