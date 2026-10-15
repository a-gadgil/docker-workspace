use std::io::{Read, Write};
use std::net::{TcpListener, TcpStream};
use std::thread;

fn handle_client(mut stream: TcpStream) {
    let mut buffer = [0; 1024];
    let _ = stream.read(&mut buffer);

    let response = "HTTP/1.1 200 OK\r\nContent-Type: application/json\r\n\r\n{\"status\":\"healthy\",\"language\":\"rust\",\"tier\":\"micro-service\"}";
    let _ = stream.write_all(response.as_bytes());
    let _ = stream.flush();
}

fn main() {
    let listener = TcpListener::bind("0.0.0.0:8081").unwrap();
    println!("Rust micro-service listening on port 8081...");

    for stream in listener.incoming() {
        match stream {
            Ok(stream) => {
                thread::spawn(|| {
                    handle_client(stream);
                });
            }
            Err(e) => {
                eprintln!("Connection failed: {}", e);
            }
        }
    }
}