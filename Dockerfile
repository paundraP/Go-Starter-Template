# Build Stage
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the main application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main ./cmd/main.go

# Final Minimal Image
FROM alpine:latest

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose the port your app runs on
EXPOSE 8080

# Run the binary
CMD ["/root/main"]
