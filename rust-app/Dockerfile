# --- Stage 1: Build Stage ---
FROM rust:alpine AS builder
WORKDIR /app
COPY . .
# Build a fully static release binary
RUN cargo build --release

# --- Stage 2: Runtime Stage ---
FROM alpine:3.19
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/target/release/rust-app /app/rust-app

EXPOSE 8080
CMD ["./rust-app"]