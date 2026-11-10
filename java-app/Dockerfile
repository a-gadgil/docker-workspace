# --- Stage 1: Build the Java Artifact ---
FROM maven:3.9-eclipse-temurin-21 AS builder

WORKDIR /app

COPY pom.xml .
RUN mvn dependency:go-offline
COPY src ./src
RUN mvn package -DskipTests

# --- Stage 2: Secure Runtime Image ---
FROM eclipse-temurin:21-jre-alpine
WORKDIR /app

# Create a non-root system user and group
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

# Copy the packaged jar from the builder stage
COPY --from=builder /app/target/java-app-1.0-SNAPSHOT.jar /app/app.jar

# Change ownership to non-root user
RUN chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
