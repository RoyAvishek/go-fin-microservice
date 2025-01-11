# Build stage
FROM golang:1.21.3-alpine AS builder

# Install necessary build tools
RUN apk add --no-cache git

WORKDIR /app

# Copy only dependency files first
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s' \
    -o main .

# Final stage
FROM alpine:3.19

WORKDIR /app

# Copy only the binary from builder
COPY --from=builder /app/main .

# Copy config file
COPY --from=builder /app/internal/config/config.yaml ./internal/config/config.yaml

# Add necessary certificates for HTTPS
RUN apk --no-cache add ca-certificates

EXPOSE 8080

CMD ["./main"]