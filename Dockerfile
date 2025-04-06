# Use a Go base image
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install required build dependencies (for SQLite)
RUN apk add --no-cache \
    build-base \
    sqlite-dev

# Copy the Go modules and install dependencies
COPY server/go.mod server/go.sum ./
RUN go mod download

# Copy the application source code
COPY server/. .

# Set CGO_ENABLED=1 and build the Go binary with sqlite support
ENV CGO_ENABLED=1

# Build the Go application
RUN go build -o rest-api-pod .

# Use a minimal base image for the final image
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache sqlite-libs

# Set the working directory inside the container
WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/rest-api-pod .

# Expose the port the application will listen on
EXPOSE 8080

# Run the binary when the container starts
CMD ["./rest-api-pod"]

