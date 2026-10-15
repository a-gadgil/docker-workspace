# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Install git and dependencies if needed for private modules
RUN apk add --no-cache git

# Set working directory inside the container
WORKDIR /app

# Copy dependency manifests first to leverage Docker layer caching
COPY go.mod go.sum* ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build a statically linked binary (CGO_ENABLED=0 removes C dependency)
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/main .

# Stage 2: Runtime image
FROM alpine:3.19

# Add CA certificates for secure connections (HTTPS/TLS to Elasticsearch/MongoDB)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/bin/main /app/main

# Expose port 8080 to match Docker Compose port mappings
EXPOSE 8080

# Run the application binary
CMD ["./main"]