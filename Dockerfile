# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

# Set workdir
WORKDIR /app

# Cache Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source (excludes vendor/ for layer efficiency)
COPY . .

# Build with optimizations: static, trim paths, strip debug
RUN GOOS=linux GOARCH=amd64 go build \
    -ldflags="-w -s -buildid=" -trimpath \
    -o annora-lore ./cmd/annora-lore

# ---------- Runtime stage ----------
FROM gcr.io/distroless/static-debian12:nonroot

WORKDIR /app

LABEL org.opencontainers.image.source="https://github.com/AdityaTaggar05/annora-lore"
LABEL org.opencontainers.image.description="Lore Service: Graph-based lore management"

COPY --from=builder /app/annora-lore .
COPY ./internal/repository/queries ./internal/repository/queries

EXPOSE 8000

# Health check (assumes /health endpoint from your main.go)
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8000/health || exit 1

CMD ["./annora-lore"]
