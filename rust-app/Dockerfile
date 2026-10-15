# Stage 1: Build the Rust binary
FROM rust:1.75-alpine AS builder

WORKDIR /app

# Copy dependency manifests
COPY Cargo.toml ./

# Create a dummy src/main.rs to build dependencies and cache them
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo build --release

# Copy actual source code
COPY src ./src

# Touch main.rs to ensure the compiler recognizes source changes and builds the real binary
RUN touch src/main.rs
RUN cargo build --release

# Stage 2: Runtime image
FROM alpine:3.19

# Add CA certificates for secure network calls if needed
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the compiled release binary from the builder stage
COPY --from=builder /app/target/release/rust-app /app/rust-app

# Expose port 8081 to match the TCP listener in main.rs
EXPOSE 8081

# Run the binary
CMD ["./rust-app"]