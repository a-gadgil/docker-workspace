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

# Command to run the application from the src directory
CMD ["node", "src/index.js"]