# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# CGO_ENABLED=0 for a static binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /api ./cmd/api

# Final stage - Distroless for security (no shell, no package manager)
FROM gcr.io/distroless/static-debian12:latest

WORKDIR /

COPY --from=builder /api /api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/api"]
