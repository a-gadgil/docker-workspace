# --- Stage 1: Build Go Binary ---
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Cache Go modules first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-app .

# --- Stage 2: Minimal Non-Root Runtime ---
FROM alpine:3.20

WORKDIR /app

# Install ca-certificates and create non-root user while still root
RUN apk add --no-cache ca-certificates && \
    addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

COPY --from=builder /app/go-app /app/go-app

RUN chown -R appuser:appgroup /app

# Switch to non-root user before starting the service
USER appuser

EXPOSE 8080

CMD ["/app/go-app"]
