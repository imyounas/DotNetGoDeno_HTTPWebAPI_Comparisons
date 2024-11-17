# Build stage
FROM golang:1.23-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go mod file
COPY go.mod ./

# Download dependencies (this will create go.sum)
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/main .

# Expose port (adjust as needed for your application)
EXPOSE 8080

# Run the binary
CMD ["./main"]