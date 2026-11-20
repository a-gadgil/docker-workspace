#include <csignal>
#include <cstdlib>
#include <cstring>
#include <iostream>
#include <netinet/in.h>
#include <string>
#include <sys/socket.h>
#include <unistd.h>

namespace {
volatile sig_atomic_t g_running = 1;
int g_server_fd = -1;

void signal_handler(int) {
    g_running = 0;
    if (g_server_fd >= 0) {
        shutdown(g_server_fd, SHUT_RDWR);
    }
}

std::string get_env_or_default(const char* name, const std::string& fallback) {
    const char* val = std::getenv(name);
    if (val == nullptr || val[0] == '\0') {
        return fallback;
    }
    return std::string(val);
}

int parse_port(const std::string& port_str, int fallback) {
    try {
        int parsed = std::stoi(port_str);
        if (parsed > 0 && parsed <= 65535) {
            return parsed;
        }
    } catch (...) {
        // Fall back below
    }
    return fallback;
}

std::string build_http_response(const std::string& content_type, const std::string& body) {
    return "HTTP/1.1 200 OK\r\n"
           "Content-Type: " + content_type + "\r\n"
           "Content-Length: " + std::to_string(body.size()) + "\r\n"
           "Connection: close\r\n\r\n" + body;
}
}  // namespace

int main() {
    std::signal(SIGINT, signal_handler);
    std::signal(SIGTERM, signal_handler);

    const std::string env_name = get_env_or_default("ENVIRONMENT", "development");
    const std::string service_name = get_env_or_default("SERVICE_NAME", "cpp-secure-app");
    const int port = parse_port(get_env_or_default("PORT", "8080"), 8080);

    g_server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (g_server_fd < 0) {
        std::cerr << "[ERROR] Failed to create socket." << std::endl;
        return 1;
    }

    int opt = 1;
    setsockopt(g_server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

    sockaddr_in address{};
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(static_cast<uint16_t>(port));

    if (bind(g_server_fd, reinterpret_cast<sockaddr*>(&address), sizeof(address)) < 0) {
        std::cerr << "[ERROR] Failed to bind to port " << port << std::endl;
        close(g_server_fd);
        return 1;
    }

    if (listen(g_server_fd, 16) < 0) {
        std::cerr << "[ERROR] Failed to listen on socket." << std::endl;
        close(g_server_fd);
        return 1;
    }

    std::cout << "[" << service_name << "] listening on 0.0.0.0:" << port
              << " (" << env_name << ")" << std::endl;

    while (g_running) {
        sockaddr_in client_addr{};
        socklen_t client_len = sizeof(client_addr);
        int client_fd = accept(g_server_fd, reinterpret_cast<sockaddr*>(&client_addr), &client_len);
        if (client_fd < 0) {
            if (!g_running) {
                break;
            }
            continue;
        }

        char buffer[2048]{};
        ssize_t bytes_read = read(client_fd, buffer, sizeof(buffer) - 1);
        if (bytes_read > 0) {
            std::string request(buffer, static_cast<size_t>(bytes_read));
            std::string response;

            if (request.find("GET /health") == 0) {
                std::string json_body = "{\"status\":\"healthy\",\"environment\":\"" +
                                        env_name + "\",\"service\":\"" + service_name + "\"}";
                response = build_http_response("application/json; charset=utf-8", json_body);
            } else {
                std::string text_body = "Hello from " + service_name + "!\n";
                response = build_http_response("text/plain; charset=utf-8", text_body);
            }

            write(client_fd, response.data(), response.size());
        }

        close(client_fd);
    }

    if (g_server_fd >= 0) {
        close(g_server_fd);
        g_server_fd = -1;
    }

    std::cout << "[" << service_name << "] shutdown complete." << std::endl;
    return 0;
}
