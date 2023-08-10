# Use an official Go runtime as a base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's workspace
COPY . .

# Build the Go application
RUN go build -o app .

# Expose a port on which your application will listen
EXPOSE 8080

# Command to run the executable
CMD ["./app"]
