# 1. Base image layer
FROM golang:1.22-alpine AS builder
WORKDIR /app

# 2. Dependency layer (Cached unless go.mod / go.sum change)
COPY go.mod ./
RUN go mod download

# 3. Source code layer (Rebuilds only when main.go changes)
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 4. Minimal runtime layer
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]