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
# Stage 2: Production Runtime Stage
# ==========================================
FROM alpine:3.19

WORKDIR /root/

# Install CA certificates for outbound TLS connections if required
RUN apk --no-cache add ca-certificates

# Security Best Practice: Run as a non-root user
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -D appuser

USER appuser

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose the application port
EXPOSE 8080

# Command to run the application
CMD ["./server"]
