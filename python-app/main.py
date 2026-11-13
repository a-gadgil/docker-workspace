import os
from flask import Flask, jsonify

app = Flask(__name__)

@app.route('/')
def hello():
    return 'Hello from Python Microservice running on Linux!\n'

@app.route('/health', methods=['GET'])
def health():
    env = os.getenv('ENVIRONMENT', 'development')
    return jsonify({
        "status": "healthy",
        "environment": env,
        "service": "python-secure-app"
    }), 200

if __name__ == '__main__':
    port = int(os.getenv('PORT', 8080))
    app.run(host='0.0.0.0', port=port)
