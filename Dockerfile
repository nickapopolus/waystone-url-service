FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Copy source code
COPY . .
RUN if [ -f "../go.sum" ]; then go mod download; fi || true

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o url-service ./cmd/server

# Final stage
FROM alpine:latest

# Add ca-certificates for HTTPS calls
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /app/url-service .

# Expose port
EXPOSE 50001

# Run
CMD ["./url-service"]