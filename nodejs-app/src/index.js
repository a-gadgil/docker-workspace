const express = require('express');
const { Pool } = require('pg');
const { createClient } = require('redis');

const app = express();
const port = process.env.PORT || 3000;

// PostgreSQL Connection Pool using Compose service name 'postgres'
const pool = new Pool({
  host: process.env.DB_HOST || 'postgres',
  port: process.env.DB_PORT || 5432,
  user: process.env.DB_USER || 'app_user',
  password: process.env.DB_PASSWORD || 'secret_password',
  database: process.env.DB_NAME || 'app_db',
});

// Redis Client using Compose service name 'redis'
const redisClient = createClient({
  url: `redis://${process.env.REDIS_HOST || 'redis'}:${process.env.REDIS_PORT || 6379}`
});

redisClient.on('error', (err) => console.error('Redis Client Error', err));

// Initialize connections and start server
async function startApp() {
  try {
    // Connect to Redis
    await redisClient.connect();
    console.log('Connected to Redis successfully');

    // Test Postgres Connection & create a sample table if it doesn't exist
    await pool.query(`
      CREATE TABLE IF NOT EXISTS visits (
        id SERIAL PRIMARY KEY,
        count INT NOT NULL,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
      );
    `);
    console.log('Connected to PostgreSQL successfully');

    app.get('/', async (req, res) => {
      try {
        // Increment visit counter in Redis cache
        const cachedVisits = await redisClient.incr('visits_count');

        // Log or fetch from Postgres for demonstration
        const dbResult = await pool.query('SELECT NOW()');

        res.json({
          message: 'Welcome to the Docker Practice Lab!',
          redis_visits_counter: cachedVisits,
          postgres_time: dbResult.rows[0].now,
        });
      } catch (err) {
        console.error(err);
        res.status(500).json({ error: 'Internal Server Error' });
      }
    });

    app.listen(port, () => {
      console.log(`App running on port ${port}`);
    });

  } catch (err) {
    console.error('Failed to start application:', err);
    process.exit(1);
  }
}

startApp();