# Stage 1: Build the binary with cached dependencies
FROM rust:alpine AS builder

# Install build tools AND static OpenSSL libraries required for musl static linking
RUN apk add --no-cache pkgconfig openssl-dev openssl-libs-static build-base

WORKDIR /app

# 1. Copy Cargo.toml and optional Cargo.lock safely
COPY Cargo.toml Cargo.lock* ./

# 2. Create a dummy main.rs so Cargo can download and build dependencies without your source code yet
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo build --release

# 3. Now copy your actual source code over the dummy src
COPY src ./src

# 4. Touch main.rs to force Cargo to recognize the real code change, then build
RUN touch src/main.rs
RUN cargo build --release

# Stage 2: Runtime image (Ultra-lightweight)
FROM alpine:3.19
RUN apk add --no-cache openssl ca-certificates
WORKDIR /app
COPY --from=builder /app/target/release/rust_app /app/rust_app

EXPOSE 8080
CMD ["./rust_app"]