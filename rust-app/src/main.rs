use std::env;
use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};

fn handle_client(mut stream: TcpStream) {
    let mut buffer = [0; 1024];
    if stream.read(&mut buffer).is_err() {
        return;
    }

    let request = String::from_utf8_lossy(&buffer[..]);

    // Check for health endpoint
    if request.starts_with("GET /health") {
        let env = env::var("ENVIRONMENT").unwrap_or_else(|_| "development".to_string());
        let json_response = format!(
            "{{\"status\":\"healthy\",\"environment\":\"{}\",\"service\":\"rust-secure-app\"}}",
            env
        );

        let response = format!(
            "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\nContent-Length: {}\r\n\r\n{}",
            json_response.len(),
            json_response
        );
        let _ = stream.write_all(response.as_bytes());
    } else {
        let body = "Hello from Secure Rust Microservice!\n";
        let response = format!(
            "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: {}\r\n\r\n{}",
            body.len(),
            body
        );
        let _ = stream.write_all(response.as_bytes());
    }
    let _ = stream.flush();
}

fn main() {
    let port = env::var("PORT").unwrap_or_else(|_| "8080".to_string());
    let addr = format!("0.0.0.0:{}", port);

    let listener = TcpListener::bind(&addr).expect("Failed to bind port");
    println!("Rust secure microservice listening on {}", addr);

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                std::thread::spawn(|| handle_client(stream));
            }
            Err(e) => {
                eprintln!("Connection failed: {}", e);
            }
        }
    }
}
