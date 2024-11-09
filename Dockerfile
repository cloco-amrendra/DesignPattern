# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Install Air for live reloading
RUN curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project (including singleton and factory folders)
COPY . .

# Expose port 8080
EXPOSE 8080

# Use Air to run the app with live reloading
CMD ["air"]
