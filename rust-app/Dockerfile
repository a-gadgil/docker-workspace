# --- Stage 1: Build the Rust Binary ---
FROM rust:1.77-alpine AS builder

WORKDIR /app

# Copy dependency manifest and source code together
COPY Cargo.toml ./
COPY src ./src

# Generate lockfile and build release binary
RUN cargo build --release

# --- Stage 2: Secure Runtime Image ---
FROM alpine:3.20

WORKDIR /app

# Create a non-root system user and group (UID/GID 1000)
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

# Copy the compiled binary from the builder stage
COPY --from=builder /app/target/release/rust-app /app/rust-app

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 8080

CMD ["/app/rust-app"]
