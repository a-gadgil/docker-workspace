# ==========================================
# Build Stage
# ==========================================

FROM golang:1.26-alpine AS builder

WORKDIR /app

# Install dependencies

COPY go.mod ./
RUN go mod download

# Copy source code and build statically linked binary

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go


# ==========================================
# Runtime Stage
# ==========================================
FROM alpine:3.19

WORKDIR /app

FROM alpine:3.21

# Security Best Practice: Run as a non-root user (UID 10001 to match your Kubernetes securityContext)
RUN addgroup -g 10001 appgroup && \
    adduser -u 10001 -G appgroup -D appuser

# Copy the compiled binary from the builder stage with correct ownership
COPY --chown=appuser:appgroup --from=builder /app/server /app/server

# Ensure the binary is executable
RUN chmod +x /app/server

# Switch to non-root user
USER appuser
RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

# Command to run the application using absolute path
CMD ["/app/server"]
USER 10001:10001

CMD ["./server"]
