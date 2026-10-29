#include "httplib.h"
#include <libpq-fe.h>
#include <hiredis/hiredis.h>
#include <iostream>
#include <sstream>

int main() {
    httplib::Server svr;

    // Health check endpoint
    svr.Get("/health", [](const httplib::Request& req, httplib::Response& res) {
        res.set_content("OK", "text/plain");
    });

    // Root endpoint with DB & Redis verification
    svr.Get("/", [](const httplib::Request& req, httplib::Response& res) {
        // 1. Connect to PostgreSQL
        const char* conninfo = "postgresql://user:password@postgres_db:5432/mydb";
        PGconn *conn = PQconnectdb(conninfo);
        std::string db_status = "Connected successfully";
        if (PQstatus(conn) != CONNECTION_OK) {
            db_status = PQerrorMessage(conn);
        }
        PQfinish(conn);

        // 2. Connect to Redis
        redisContext *c = redisConnect("redis_cache", 6379);
        std::string redis_status = "Connected successfully";
        if (c == NULL || c->err) {
            if (c) {
                redis_status = c->errstr;
                redisFree(c);
            } else {
                redis_status = "Connection error: can't allocate redis context";
            }
        } else {
            redisReply *reply = (redisReply*)redisCommand(c, "INCR visitor_count");
            if (reply) {
                freeReplyObject(reply);
            }
            redisFree(c);
        }

        // Build JSON response
        std::stringstream json;
        json << "{"
             << "\"status\":\"success\","
             << "\"message\":\"Hello from C++ production container!\","
             << "\"database\":\"" << db_status << "\","
             << "\"redis\":\"" << redis_status << "\""
             << "}";

        res.set_content(json.str(), "application/json");
    });

    std::cout << "C++ HTTP server starting on port 8080..." << std::endl;
    svr.listen("0.0.0.0", 8080);
    return 0;
}
