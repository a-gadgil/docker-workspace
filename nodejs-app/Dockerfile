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

EXPOSE 3000

HEALTHCHECK --interval=30s --timeout=3s CMD wget --spider http://localhost:3000/health || exit 1

CMD ["node", "src/index.js"]