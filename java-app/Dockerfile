# --- Stage 1: Build the Java application with Maven ---
FROM maven:3.9-eclipse-temurin-17-alpine AS builder

WORKDIR /app

# Copy pom.xml and source code
COPY pom.xml ./
COPY src ./src

# Package the application into an executable jar
RUN mvn clean package -DskipTests

# --- Stage 2: Secure Runtime Image ---
FROM eclipse-temurin:17-jre-alpine

WORKDIR /app

# Copy the built jar from the builder stage
COPY --from=builder /app/target/java-app-1.0.0-jar-with-dependencies.jar /app/app.jar

# Create a non-root system user and group (UID/GID 1000)
RUN addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser && \
    chown -R appuser:appgroup /app

# Switch to non-root user
USER appuser

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "/app/app.jar"]
