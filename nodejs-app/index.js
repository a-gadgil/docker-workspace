const http = require('http');

const PORT = Number(process.env.PORT) || 8080;
const ENVIRONMENT = process.env.ENVIRONMENT || 'development';
const SERVICE_NAME = process.env.SERVICE_NAME || 'nodejs-secure-app';

const server = http.createServer((req, res) => {
  if (req.url === '/health' && req.method === 'GET') {
    const payload = JSON.stringify({
      status: 'healthy',
      environment: ENVIRONMENT,
      service: SERVICE_NAME,
    });
    res.writeHead(200, {
      'Content-Type': 'application/json; charset=utf-8',
      'Content-Length': Buffer.byteLength(payload),
    });
    res.end(payload);
    return;
  }

  const body = `Hello from ${SERVICE_NAME}!\n`;
  res.writeHead(200, {
    'Content-Type': 'text/plain; charset=utf-8',
    'Content-Length': Buffer.byteLength(body),
  });
  res.end(body);
});

server.listen(PORT, '0.0.0.0', () => {
  console.log(`[${SERVICE_NAME}] listening on 0.0.0.0:${PORT} (${ENVIRONMENT})`);
});

function shutdown(signal) {
  console.log(`[${SERVICE_NAME}] received ${signal}, closing HTTP server...`);
  server.close(() => {
    console.log(`[${SERVICE_NAME}] shutdown complete`);
    process.exit(0);
  });
}

process.on('SIGTERM', () => shutdown('SIGTERM'));
process.on('SIGINT', () => shutdown('SIGINT'));
