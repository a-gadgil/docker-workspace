# Stage 1: Build stage using Maven or direct javac compilation
FROM eclipse-temurin:21-jdk-alpine AS builder
WORKDIR /app
COPY Main.java .
RUN javac Main.java

# Stage 2: Runtime stage
FROM eclipse-temurin:21-jre-alpine
WORKDIR /app

# Copy compiled class file from builder stage
COPY --from=builder /app/Main.class .

EXPOSE 8080

ENV PORT=8080 \
    APP_VERSION=v2.0.0

CMD ["java", "Main"]