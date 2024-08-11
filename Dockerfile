# Use Go 1.22 as the base image
FROM golang:1.22

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o unicomer_challenge

# Expose the application port (optional, if your service listens on a specific port)
EXPOSE 8080

# Command to run the executable
CMD ["./unicomer_challenge"]
