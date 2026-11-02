# ==========================================
# Stage 1: Build Stage
# ==========================================
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install git or necessary certificate packages if needed
RUN apk add --no-cache git

# Copy dependency manifests first to leverage Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY main.go ./

# Build a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

# ==========================================
# Runtime Stage
# ==========================================
FROM alpine:3.19

WORKDIR /app

# Install CA certificates for outbound TLS connections if required
RUN apk --no-cache add ca-certificates

# Security Best Practice: Run as a non-root user (UID 10001 to match your Kubernetes securityContext)
RUN addgroup -g 10001 appgroup && \
    adduser -u 10001 -G appgroup -D appuser

# Copy the compiled binary from the builder stage with correct ownership
COPY --chown=appuser:appgroup --from=builder /app/server /app/server

# Ensure the binary is executable
RUN chmod +x /app/server

# Switch to non-root user
USER appuser

# Expose the application port
EXPOSE 8080

# Command to run the application using absolute path
CMD ["/app/server"]
