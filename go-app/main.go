package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	elastic "github.com/elastic/go-elasticsearch/v8"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Connect to MongoDB
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://admin:secretpassword@localhost:27017/appdb?authSource=admin"
	}
	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Printf("Failed to connect to MongoDB: %v", err)
	} else {
		if err := mongoClient.Ping(ctx, nil); err != nil {
			log.Printf("MongoDB ping failed: %v", err)
		} else {
			log.Println("Successfully connected to MongoDB!")
		}
	}

	// 2. Connect to Redis
	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379"
	}
	opt, err := redis.ParseURL(redisURL)
	if err != nil {
		log.Printf("Failed to parse Redis URL: %v", err)
	}
	rdb := redis.NewClient(opt)
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Printf("Redis ping failed: %v", err)
	} else {
		log.Println("Successfully connected to Redis!")
	}

	// 3. Connect to Elasticsearch
	esURL := os.Getenv("ELASTICSEARCH_URL")
	if esURL == "" {
		esURL = "http://localhost:9200"
	}
	es, err := elastic.NewClient(elastic.Config{
		Addresses: []string{esURL},
	})
	if err != nil {
		log.Printf("Failed to create Elasticsearch client: %v", err)
	} else {
		res, err := es.Info()
		if err != nil {
			log.Printf("Elasticsearch info failed: %v", err)
		} else {
			defer res.Body.Close()
			log.Println("Successfully connected to Elasticsearch!")
		}
	}

	// 4. HTTP Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "healthy", "tier": "Go API Gateway", "services": ["mongodb", "redis", "elasticsearch"]}`))
	})

	log.Println("Starting Go API server on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
