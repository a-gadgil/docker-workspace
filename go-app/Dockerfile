# --- Stage 1: Build the Go Binary ---
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Install git and dependencies if required
RUN apk add --no-cache git

# Copy dependency manifests
COPY go.mod ./
RUN go mod download

# Copy app source code
COPY . .

# Build a statically linked Go binary with CGO disabled
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

# --- Stage 2: Secure Runtime Image ---
FROM alpine:3.20

WORKDIR /app

# Create a non-root system user and group
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server /app/server

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Install CA certificates for outbound TLS connections if required
RUN apk --no-cache add ca-certificates

# Security Best Practice: Run as a non-root user (UID 10001 to match your Kubernetes securityContext)
RUN addgroup -g 10001 appgroup && \
    adduser -u 10001 -G appgroup -D appuser

EXPOSE 8080

CMD ["/app/server"]
