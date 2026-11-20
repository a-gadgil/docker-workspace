# --- Stage 2: Minimal Non-Root Runtime ---
FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache libstdc++ && \
    addgroup -g 1000 appgroup && \
    adduser -u 1000 -G appgroup -s /bin/sh -D appuser

COPY --from=builder /app/build/cpp_app /app/cpp_app

RUN chown -R appuser:appgroup /app

USER appuser

EXPOSE 8080

CMD ["/app/cpp_app"]
