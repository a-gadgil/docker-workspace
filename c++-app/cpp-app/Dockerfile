# --- Stage 1: Build Stage ---
FROM debian:bookworm-slim AS builder

RUN apt-get update && apt-get install -y \
    g++ \
    cmake \
    make \
    git \
    pkg-config \
    libpq-dev \
    libhiredis-dev \
    curl \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN curl -sSL https://raw.githubusercontent.com/yhirose/cpp-httplib/master/httplib.h -o httplib.h

COPY CMakeLists.txt main.cpp ./

RUN mkdir build && cd build && cmake .. && make

# --- Stage 2: Runtime Production Stage ---
FROM debian:bookworm-slim AS runner

RUN apt-get update && apt-get install -y \
    libpq5 \
    libhiredis0.14 \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN groupadd -g 1001 cgroup && \
    useradd -u 1001 -g cgroup -s /bin/bash cgroupuser

COPY --from=builder /app/build/cpp_app /app/cpp_app

RUN chown -R cgroupuser:cgroup /app
USER cgroupuser

EXPOSE 8080

CMD ["./cpp_app"]
