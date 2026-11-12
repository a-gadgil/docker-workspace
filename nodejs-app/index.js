const http = require('http');

const PORT = process.env.PORT || 8080;
const ENVIRONMENT = process.env.ENVIRONMENT || 'development';

const server = http.createServer((req, res) => {
    if (req.url === '/health' && req.method === 'GET') {
        res.writeHead(200, { 'Content-Type': 'application/json' });
        res.end(JSON.stringify({
            status: 'healthy',
            environment: ENVIRONMENT,
            service: 'nodejs-secure-app'
        }));
    } else {
        res.writeHead(200, { 'Content-Type': 'text/plain' });
        res.end('Hello from Secure Node.js Microservice!\n');
    }
});

server.listen(PORT, () => {
    console.log(`Node.js secure microservice listening on port ${PORT}`);
});
