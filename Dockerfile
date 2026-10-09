# ==============================================================================
# Multi-Stage Production Dockerfile for Ukacha Juma Portfolio (Render Ready)
# ==============================================================================

# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git ca-certificates tzdata

# Cache Go modules
COPY go.mod ./
RUN go mod download || true

# Copy source code
COPY . .

# Build static stripped binary for Linux
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o /app/portfolio-server .

# ==============================================================================
# Final Minimal Runtime Stage
# ==============================================================================
FROM alpine:3.20 AS runner

WORKDIR /app

# Install ca-certificates and curl for health checks
RUN apk add --no-cache ca-certificates tzdata curl && \
    addgroup -S appgroup && adduser -S appuser -G appgroup

# Copy compiled binary from builder
COPY --from=builder /app/portfolio-server /app/portfolio-server

# Copy required web runtime assets
COPY --from=builder /app/templates /app/templates
COPY --from=builder /app/static /app/static

# Permissions
RUN chown -R appuser:appgroup /app
USER appuser

# Render injects $PORT (defaults to 8080)
ENV PORT=8080
EXPOSE 8080

# Container healthcheck
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:${PORT}/healthz || exit 1

# Start server
ENTRYPOINT ["/app/portfolio-server"]
