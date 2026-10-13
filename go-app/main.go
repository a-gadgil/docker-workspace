package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	dbUser := getEnv("DB_USER", "root")
	dbPass := getEnv("DB_PASSWORD", "rootpassword")
	dbHost := getEnv("DB_HOST", "mysql")
	dbName := getEnv("DB_NAME", "appdb")
	redisAddr := getEnv("REDIS_ADDR", "redis:6379")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPass, dbHost, dbName)

	// Connect to MySQL with retry mechanism
	var db *sql.DB
	var err error
	for i := 0; i < 15; i++ {
		db, err = sql.Open("mysql", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				break
			}
		}
		log.Println("Waiting for MySQL database to become ready...")
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	defer db.Close()

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Increment visitor count in Redis cache
		count, err := rdb.Incr(ctx, "visits").Result()
		if err != nil {
			http.Error(w, "Redis error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Query current time from MySQL backend
		var now time.Time
		err = db.QueryRow("SELECT NOW()").Scan(&now)
		if err != nil {
			http.Error(w, "MySQL error: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "🚀 Go App Stack Running via Docker Compose!\n")
		fmt.Fprintf(w, "📊 Redis Visit Count: %d\n", count)
		fmt.Fprintf(w, "🗄️ MySQL Database Time: %v\n", now)
	})

	log.Println("Go application server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
