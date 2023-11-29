# Use the official Golang image as the base image
FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go code into the container at /app
COPY . .

# Download Go dependencies
RUN go get -d -v ./...

# Build the Go application
RUN go build -o /bankapi .

# Use a minimal base image for the final stage
FROM alpine:latest

# Set the working directory in the final image
WORKDIR /app

# Copy only the binary from the builder stage
COPY --from=builder /bankapi .

# Expose the port on which the application will run
EXPOSE 8080

# Command to run the executable
CMD ["./bankapi"]
