# Dockerfile for development

# Use an official Go runtime as a parent image
FROM golang:1.24.2 AS development

# Set the working directory inside the container
WORKDIR /app

# Install air for live reloading
# (Ensure you have network access during build if behind proxy)
RUN go install github.com/air-verse/air@latest

# Copy go mod and sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
# Note: This will be overlayed by the volume mount during dev,
# but it's good practice for potential image rebuilding.
COPY . .

# Copy the air config file
COPY .air.toml .

# Expose the port the app runs on
EXPOSE $SERVER_PORT

# Command to run air, which will build and run the app
# Air will watch for changes in the mounted volume
CMD ["air"]

