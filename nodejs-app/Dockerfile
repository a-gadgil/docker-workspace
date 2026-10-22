FROM node:18-alpine

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