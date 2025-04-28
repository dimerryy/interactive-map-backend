# Use Go base image
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the app
RUN go build -o main .

# Run the app
CMD ["./main"]
