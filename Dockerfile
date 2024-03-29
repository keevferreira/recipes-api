# Use the official Golang image as a base image
FROM golang:1.21-alpine AS build

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go app
RUN go build -o recipes-api ./cmd/recipes-api

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the previous stage
COPY --from=build /app/recipes-api .

# Expose port 8080 to the outside world
EXPOSE ${SERVER_PORT}

# Command to run the executable
CMD ["./recipes-api"]
