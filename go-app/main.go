package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Read connection strings from environment variables defined in Docker Compose
	mysqlDSN := os.Getenv("MYSQL_DSN")
	redisAddr := os.Getenv("REDIS_ADDR")
	mongoURI := os.Getenv("MONGO_URI")
	esURL := os.Getenv("ES_URL")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Test MySQL Connection
	db, err := sql.Open("mysql", mysqlDSN)
	if err == nil && db.PingContext(ctx) == nil {
		log.Println("[+] Successfully connected to MySQL")
	} else {
		log.Printf("[-] MySQL connection warning: %v", err)
	}

	// 2. Test Redis Connection
	rdb := redis.NewClient(&redis.Options{Addr: redisAddr})
	if _, err := rdb.Ping(ctx).Result(); err == nil {
		log.Println("[+] Successfully connected to Redis")
	} else {
		log.Printf("[-] Redis connection warning: %v", err)
	}

	// 3. Test MongoDB Connection
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err == nil && mongoClient.Ping(ctx, nil) == nil {
		log.Println("[+] Successfully connected to MongoDB")
	} else {
		log.Printf("[-] MongoDB connection warning: %v", err)
	}

	// 4. Log Elasticsearch Endpoint Status
	log.Printf("[*] Elasticsearch target endpoint configured at: %s", esURL)

	// Simple HTTP Health & Status Server
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("5-Tier App Stack is Healthy and Running!"))
	})

	log.Println("[*] Go application server starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
