# Use the official Go image as a base image
FROM golang:1.20

# Set the working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Expose the application port
EXPOSE 3000

# Run the application
CMD ["./main"]
