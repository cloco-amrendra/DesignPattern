# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the root go.mod and go.sum for dependency installation
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project (including singleton and factory folders)
COPY . .

# Expose port 8080
EXPOSE 8080
