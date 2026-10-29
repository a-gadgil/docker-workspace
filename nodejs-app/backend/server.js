const express = require('express');
const { Pool } = require('pg');
const { createClient } = require('redis');

const app = express();
const PORT = process.env.PORT || 4000;

// PostgreSQL Connection
const pool = new Pool({
    connectionString: `postgresql://appuser:secretpassword@database:5432/appdb`
});

// Redis Connection
const redisClient = createClient({
    url: 'redis://cache:6379'
});

redisClient.connect().catch(console.error);

app.get('/api/data', async (req, res) => {
    try {
        // Test DB query
        const dbResult = await pool.query('SELECT NOW()');

        // Test Redis cache increment
        const visits = await redisClient.incr('visits');

        res.json({
            message: 'Hello from Backend API!',
            dbTime: dbResult.rows[0].now,
            visitCount: visits
        });
    } catch (err) {
        console.error(err);
        res.status(500).json({ error: 'Internal server error' });
    }
});

app.listen(PORT, () => {
    console.log(`Backend running on port ${PORT}`);
});
