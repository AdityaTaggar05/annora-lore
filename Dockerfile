# syntax=docker/dockerfile:1  # Enables BuildKit features like cache mounts (optional for 2026+)

# Stage 1: Build the Go binary
FROM golang:1.22-alpine AS builder

# Install git (for go mod deps; remove if not needed)
RUN apk add --no-cache git

# Set workdir
WORKDIR /app

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source (excludes vendor/ for layer efficiency)
COPY . .

# Build with optimizations: static, trim paths, strip debug
ARG CGO_ENABLED=0
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build \
    -ldflags="-w -s -buildid=" \
    -trimpath \
    -o annora-lore \
    ./cmd/annora-lore

# Stage 2: Runtime (distroless for security/minimalism)
FROM gcr.io/distroless/static-debian12:latest

# Labels for metadata (best practice for prod)
LABEL org.opencontainers.image.source="https://github.com/AdityaTaggar05/annora-lore"
LABEL org.opencontainers.image.description="Lore Service: Graph-based lore management"

# Copy binary from builder
COPY --from=builder /app/annora-lore /

# Non-root user (distroless default: 65532)
USER 65532:65532

# Expose port
EXPOSE 8080

# Health check (assumes /health endpoint from your main.go)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

# Workdir for consistency
WORKDIR /

# Run the binary
ENTRYPOINT ["/annora-lore"]