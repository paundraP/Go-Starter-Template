# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy only go.mod and go.sum first to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy all source files and build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Final Image (Minimal)
FROM alpine:latest

WORKDIR /root/

# Copy binary from builder stage
COPY --from=builder /docker-gs-ping .

EXPOSE 8080

# Run the binary
CMD ["/root/docker-gs-ping"]
