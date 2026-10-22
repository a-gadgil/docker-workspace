const express = require('express');
const { Pool } = require('pg');

const app = express();
const PORT = process.env.PORT || 3000;
const APP_VERSION = process.env.APP_VERSION || 'v1.0.1';

// Database connection pool using environment variables
const pool = new Pool({
  host: process.env.DB_HOST || 'localhost',
  port: process.env.DB_PORT || 5432,
  user: process.env.DB_USER || 'postgres',
  password: process.env.DB_PASSWORD || 'postgres',
  database: process.env.DB_NAME || 'devops_db',
});

// Middleware to parse JSON
app.use(express.json());

// Root Endpoint
app.get('/', async (req, res) => {
  try {
    const dbResult = await pool.query('SELECT NOW()');
    res.json({
      status: 'running',
      message: 'Node.js connected to PostgreSQL database successfully!',
      version: APP_VERSION,
      dbTime: dbResult.rows[0].now,
    });
  } catch (err) {
    res.status(500).json({
      status: 'error',
      message: 'Failed to establish connection to PostgreSQL database',
      version: APP_VERSION,
      error: err.message,
    });
  }
});

// Health Check Endpoint (useful for Docker/Kubernetes health checks)
app.get('/health', async (req, res) => {
  try {
    await pool.query('SELECT 1');
    res.status(200).json({ status: 'UP', database: 'connected' });
  } catch (err) {
    res.status(503).json({ status: 'DOWN', database: 'disconnected', error: err.message });
  }
});

// Start Server
const server = app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT} (Version: ${APP_VERSION})`);
});

// Graceful Shutdown
process.on('SIGTERM', () => {
  console.log('SIGTERM signal received: closing HTTP server and DB pool');
  server.close(async () => {
    await pool.end();
    console.log('HTTP server and DB pool closed');
    process.exit(0);
  });
});