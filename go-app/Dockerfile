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

FROM alpine:3.21

RUN apk add --no-cache ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .

EXPOSE 8080

USER 10001:10001

CMD ["./server"]
