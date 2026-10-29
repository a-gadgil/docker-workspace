use axum::{routing::get, Router};
use sqlx::MySqlPool;
use std::env;

#[tokio::main]
async fn main() {
    // 1. Read environment variables defined in Docker Compose
    let database_url = env::var("DATABASE_URL")
        .unwrap_or_else(|_| "mysql://root:secretpassword@localhost:3306/app_db".to_string());
    let memcached_url = env::var("MEMCACHED_URL")
        .unwrap_or_else(|_| "tcp://localhost:11211".to_string());
    // 2. Connect to MySQL Pool
    println!("Connecting to MySQL...");
    let db_pool = MySqlPool::connect(&database_url)
        .await
        .expect("Failed to connect to MySQL");
    println!("Successfully connected to MySQL!");

    // 3. Connect to Memcached
    println!("Connecting to Memcached...");
    let mut cache_client = memcache::Client::connect(memcached_url)
        .expect("Failed to connect to Memcached");

    // Test write/read to Memcached (Fixed with [..])
    cache_client.set("startup_key", &b"Hello from Rust & Memcached!"[..], 60).unwrap();
    let val: String = cache_client.get("startup_key").unwrap().unwrap();
    println!("Memcached test read: {}", val);

    // 4. Build Axum Router
    let app = Router::new().route("/", get(|| async { "Rust 3-Tier App is running with MySQL & Memcached!" }));

    // 5. Bind and Listen
    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await.unwrap();
    println!("Server listening on port 8080...");
    axum::serve(listener, app).await.unwrap();
}
