import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpExchange;
import java.net.InetSocketAddress;
import java.io.OutputStream;
import java.io.IOException;
import java.net.InetAddress;

public class Main {
    public static void main(String[] args) throws Exception {
        // Read port from environment variable or default to 8080
        String portEnv = System.getenv("PORT");
        int port = (portEnv != null && !portEnv.isEmpty()) ? Integer.parseInt(portEnv) : 8080;

        // Read app version from environment variable or default to v2.0.0
        String appVersion = System.getenv("APP_VERSION");
        if (appVersion == null || appVersion.isEmpty()) {
            appVersion = "v2.0.0";
        }

        String hostname = InetAddress.getLocalHost().getHostName();

        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);

        // Root Context ("/") - Returns structured JSON
        final String version = appVersion;
        server.createContext("/", exchange -> {
            String jsonResponse = String.format(
                "{\"status\":\"success\",\"message\":\"Java application running inside Docker!\",\"version\":\"%s\",\"hostname\":\"%s\"}\n",
                version, hostname
            );
            sendJsonResponse(exchange, 200, jsonResponse);
        });

        // Health Check Endpoint ("/health")
        server.createContext("/health", exchange -> {
            String jsonResponse = "{\"status\":\"UP\"}\n";
            sendJsonResponse(exchange, 200, jsonResponse);
        });

        // Info Endpoint ("/info")
        server.createContext("/info", exchange -> {
            String jsonResponse = String.format(
                "{\"runtime\":\"Java %s\",\"os\":\"%s %s\",\"hostname\":\"%s\"}\n",
                System.getProperty("java.version"),
                System.getProperty("os.name"),
                System.getProperty("os.version"),
                hostname
            );
            sendJsonResponse(exchange, 200, jsonResponse);
        });

        System.out.println("Java Lab Server listening on port " + port + " (Version: " + version + ")...");
        server.setExecutor(null); // creates a default executor
        server.start();
    }

    private static void sendJsonResponse(HttpExchange exchange, int statusCode, String response) throws IOException {
        exchange.getResponseHeaders().set("Content-Type", "application/json");
        byte[] bytes = response.getBytes("UTF-8");
        exchange.sendResponseHeaders(statusCode, bytes.length);
        OutputStream os = exchange.getResponseBody();
        os.write(bytes);
        os.close();
    }
}