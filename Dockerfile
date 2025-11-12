# Use the official Go image
FROM golang:1.25.1-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum first (for dependency caching)
COPY go.mod go.sum ./
RUN go mod download

# Download dependencies
RUN go mod download

# Copy the rest of your source code
COPY . .

# Build the Go app binary
RUN go build -o email_service .

#Expose Port for health check
EXPOSE 8080

# Tell Docker how to start your app
CMD ["./email_service"]
