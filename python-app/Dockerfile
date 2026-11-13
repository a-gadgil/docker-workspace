# --- Stage 1: Build & Install Dependencies ---
FROM python:3.12-slim AS builder

WORKDIR /app

# Install build essentials if needed
RUN apt-get update && apt-get install -y --no-install-recommends gcc && rm -rf /var/lib/apt/lists/*

# Copy and install dependencies into a local prefix directory
COPY requirements.txt .
RUN pip install --no-cache-dir --prefix=/install -r requirements.txt

# Copy source code
COPY main.py .

# --- Stage 2: Secure Runtime Image ---
FROM python:3.12-slim AS runner

WORKDIR /app

# Create a non-root system user and group (UID/GID 1000)
RUN groupadd -g 1000 appgroup && \
    useradd -u 1000 -g appgroup -s /bin/sh -m appuser

# Copy installed site-packages from builder stage
COPY --from=builder /install /usr/local

# Copy application source code
COPY --from=builder /app/main.py /app/main.py

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 8080

CMD ["gunicorn", "--bind", "0.0.0.0:8080", "main:app"]
