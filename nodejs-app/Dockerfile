# --- Secure Single-Stage or Streamlined Runtime Image ---
FROM node:20-alpine

WORKDIR /app

# Create a non-root system user (node user is built into node:20-alpine)
USER node

# Copy application files and set ownership directly
COPY --chown=node:node package.json index.js ./

EXPOSE 8080

CMD ["node", "index.js"]
