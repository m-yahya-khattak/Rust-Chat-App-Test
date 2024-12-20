# Use the official Go image for building the application
FROM golang:1.20 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest and install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the application code
COPY . .

# Build the Go application
RUN go build -o chat-app .

# Use a compatible runtime image
FROM ubuntu:22.04

# Set the working directory for the runtime container
WORKDIR /app

# Install necessary runtime dependencies
RUN apt-get update && apt-get install -y libc6 libstdc++6

# Copy the compiled application from the builder stage
COPY --from=builder /app/chat-app .

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./chat-app"]
