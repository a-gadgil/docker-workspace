const express = require('express');
const axios = require('axios');
const app = express();
const PORT = process.env.PORT || 3000;
const API_URL = process.env.API_URL || 'http://backend:4000';

app.get('/', async (req, res) => {
    try {
        const response = await axios.get(`${API_URL}/api/data`);
        res.send(`<h1>Frontend Service</h1><p>Backend says: ${response.data.message}</p>`);
    } catch (error) {
        res.send(`<h1>Frontend Service</h1><p style="color:red;">Could not connect to backend API.</p>`);
    }
});

app.listen(PORT, () => {
    console.log(`Frontend running on port ${PORT}`);
});
