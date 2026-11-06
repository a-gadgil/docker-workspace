const http = require('http');

const PORT = process.env.PORT || 8080;

const server = http.createServer((req, res) => {
    if (req.url === '/health' && req.method === 'GET') {
        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({ status: 'UP', service: 'nodejs-app' }));
    } else {
        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end('Hello from Refactored Node.js Microservice!\n');
    }
});

server.listen(PORT, () => {
    console.log(`Node.js microservice listening on port ${PORT}`);
});
