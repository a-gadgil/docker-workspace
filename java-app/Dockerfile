# --- Stage 1: Build the Java application using Gradle ---
FROM gradle:8.14-jdk21-alpine AS builder
WORKDIR /app
COPY build.gradle settings.gradle ./
COPY src ./src
RUN gradle jar --no-daemon

# --- Stage 2: Runtime Production Image ---
FROM eclipse-temurin:21-jre-alpine AS runner
WORKDIR /app

# Create non-root system user for security
RUN addgroup -g 1001 javagroup && \
    adduser -u 1001 -G javagroup -S javauser

# Copy the built jar from the builder stage
COPY --from=builder /app/build/libs/*.jar app.jar

USER javauser

EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s CMD wget --spider http://localhost:8080/health || exit 1

ENTRYPOINT ["java", "-jar", "app.jar"]