FROM node:18-alpine
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

# For Debian/Ubuntu-based images
RUN apk update && apk upgrade --no-cache

WORKDIR /app

# Copy dependency manifests first
COPY package*.json ./

# Cache layer: This step re-runs ONLY if package.json or package-lock.json changes
RUN npm ci --only=production

# Copy the rest of the application code
# Cache layer: Modifying files in src/ will break cache starting FROM THIS POINT onwards
COPY . .

# Expose the application port
EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=3s CMD wget --spider http://localhost:3000/health || exit 1

# Command to run the application from the src directory
CMD ["node", "src/index.js"]