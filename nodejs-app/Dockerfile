# --- Stage 1: Build & Install Dependencies ---
FROM node:20-alpine AS builder
WORKDIR /app
COPY package.json package-lock.json* ./
RUN npm ci --omit=dev

# --- Stage 2: Runtime Production Image ---
FROM node:20-alpine AS runner
WORKDIR /app

# Create a non-root user for security (fixing Trivy warnings!)
RUN addgroup -g 1001 nodegroup && \
    adduser -u 1001 -G nodegroup -S nodeuser

COPY --from=builder /app/node_modules ./node_modules
COPY src ./src
COPY package.json ./

USER nodeuser
# Use lightweight Node.js image
FROM node:20-alpine

# Set working directory inside container
WORKDIR /usr/src/app

# Copy dependency manifests first to leverage Docker layer caching
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application source code (including src/)
COPY . .

# Expose the application port
EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=3s CMD wget --spider http://localhost:3000/health || exit 1

# Command to run the application from the src directory
CMD ["node", "src/index.js"]