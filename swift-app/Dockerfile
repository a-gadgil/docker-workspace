# --- Stage 1: Build Stage ---
FROM swift:5.9-jammy AS builder

WORKDIR /app

# Copy package manifest and source code
COPY Package.swift ./
COPY Sources ./Sources

# Build the application in release mode with optimizations
RUN swift build -c release --static-swift-stdlib

# --- Stage 2: Runtime Production Stage ---
FROM ubuntu:22.04 AS runner

WORKDIR /app

# Create a non-root system user for security
RUN groupadd -g 1001 swiftgroup && \
    useradd -u 1001 -g swiftgroup -s /bin/bash swiftuser

# Copy the compiled binary from the builder stage
COPY --from=builder /app/.build/release/swift-app /app/swift-app

RUN chown -R swiftuser:swiftgroup /app
USER swiftuser

CMD ["./swift-app"]
