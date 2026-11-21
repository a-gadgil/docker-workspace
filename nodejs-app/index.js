const http = require("http");

const PORT = Number(process.env.PORT) || 8080;
const ENVIRONMENT = process.env.ENVIRONMENT || "development";
const SERVICE_NAME = process.env.SERVICE_NAME || "nodejs-secure-app";

function sendResponse(res, statusCode, contentType, payload) {
  const buffer = Buffer.isBuffer(payload)
    ? payload
    : Buffer.from(String(payload), "utf8");
  res.writeHead(statusCode, {
    "Content-Type": contentType,
    "Content-Length": buffer.byteLength,
  });
  res.end(buffer);
}

const server = http.createServer((req, res) => {
  const path = (req.url || "/").split("?")[0];

  if (path === "/health" && req.method === "GET") {
    const healthPayload = JSON.stringify({
      status: "healthy",
      environment: ENVIRONMENT,
      service: SERVICE_NAME,
    });
    sendResponse(res, 200, "application/json; charset=utf-8", healthPayload);
    return;
  }

  if (path === "/" && req.method === "GET") {
    sendResponse(
      res,
      200,
      "text/plain; charset=utf-8",
      `Hello from ${SERVICE_NAME}!\n`,
    );
    return;
  }

  sendResponse(res, 404, "text/plain; charset=utf-8", "Not Found\n");
});

server.listen(PORT, "0.0.0.0", () => {
  console.log(
    `[${SERVICE_NAME}] listening on 0.0.0.0:${PORT} (${ENVIRONMENT})`,
  );
});

let isShuttingDown = false;

function shutdown(signal) {
  if (isShuttingDown) {
    return;
  }
  isShuttingDown = true;

  console.log(`[${SERVICE_NAME}] received ${signal}, closing HTTP server...`);

  server.close((err) => {
    if (err) {
      console.error(`[${SERVICE_NAME}] error during server close:`, err);
      process.exit(1);
    }
    console.log(`[${SERVICE_NAME}] shutdown complete.`);
    process.exit(0);
  });

  if (typeof server.closeIdleConnections === "function") {
    server.closeIdleConnections();
  }

  const forceTimer = setTimeout(() => {
    console.warn(`[${SERVICE_NAME}] forcing exit after shutdown timeout`);
    process.exit(1);
  }, 5000);
  forceTimer.unref();
}

process.on("SIGTERM", () => shutdown("SIGTERM"));
process.on("SIGINT", () => shutdown("SIGINT"));
