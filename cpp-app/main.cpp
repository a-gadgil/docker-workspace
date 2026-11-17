#include <iostream>
#include <string>
#include <cstring>
#include <csignal>
#include <unistd.h>
#include <sys/socket.h>
#include <netinet/in.h>

volatile sig_atomic_t g_running = 1;

void signal_handler(int) {
    g_running = 0;
}

int main() {
    std::signal(SIGINT, signal_handler);
    std::signal(SIGTERM, signal_handler);

    int server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd == -1) {
        std::cerr << "[ERROR] Failed to create socket." << std::endl;
        return 1;
    }

    int opt = 1;
    setsockopt(server_fd, SOL_SOCKET, SO_REUSEADDR, &opt, sizeof(opt));

    sockaddr_in address{};
    address.sin_family = AF_INET;
    address.sin_addr.s_addr = INADDR_ANY;
    address.sin_port = htons(8080);

    if (bind(server_fd, (struct sockaddr*)&address, sizeof(address)) < 0) {
        std::cerr << "[ERROR] Failed to bind port 8080." << std::endl;
        close(server_fd);
        return 1;
    }

    if (listen(server_fd, 10) < 0) {
        std::cerr << "[ERROR] Listen failed." << std::endl;
        close(server_fd);
        return 1;
    }

    std::cout << "[INFO] C++ Zero-Dependency Microservice listening on 0.0.0.0:8080..." << std::endl;

    while (g_running) {
        sockaddr_in client_addr{};
        socklen_t client_len = sizeof(client_addr);
        int client_fd = accept(server_fd, (struct sockaddr*)&client_addr, &client_len);
        if (client_fd < 0) {
            if (!g_running) break;
            continue;
        }

        char buffer[1024] = {0};
        read(client_fd, buffer, sizeof(buffer) - 1);
        std::string req(buffer);

        std::string body;
        std::string status_line;

        if (req.find("GET /health") != std::string::npos) {
            body = R"({"status":"healthy","service":"cpp-app","version":"1.0.0"})";
            status_line = "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n";
        } else {
            body = "Hello from C++ Microservice inside docker-workspace!\n";
            status_line = "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\n";
        }

        std::string response = status_line + "Content-Length: " + std::to_string(body.size()) + "\r\n\r\n" + body;
        write(client_fd, response.c_str(), response.size());
        close(client_fd);
    }

    close(server_fd);
    std::cout << "[INFO] Server stopped cleanly." << std::endl;
    return 0;
}
