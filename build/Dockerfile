# Use an official Go runtime as a parent image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code from the host to the Working Directory inside the container
COPY . .

# Copy the service account key file (if needed)
COPY path/to/your/service-account-key.json /app/service-account-key.json

# Build the Go app
RUN go build -o main ./cmd/server

# Set environment variable for Google Application Credentials
ENV GOOGLE_APPLICATION_CREDENTIALS=/app/service-account-key.json

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]