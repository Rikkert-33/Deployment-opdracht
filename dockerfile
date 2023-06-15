# Start with a minimal Go base image
FROM golang:alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the container
COPY go.mod .

# Download and install the Go dependencies
RUN go mod download

# Copy the source code to the container
COPY . .

# Build the Go application
RUN go build -o webserver .

# Expose the port that the web application listens on
EXPOSE 8080

# Start the web application when the container starts
CMD ["./webserver"]
