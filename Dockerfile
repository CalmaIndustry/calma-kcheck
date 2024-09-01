# Use the official Golang image to create a build artifact.
# This image is based on Debian and contains the Go runtime.
FROM golang:1.22.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build main.go


# Expose port 2112 to the outside world
EXPOSE 2112
