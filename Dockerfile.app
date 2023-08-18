# Use a Go-based image as the base image
FROM golang:1.19.4

# Set the working directory inside the container
WORKDIR /app

# Install required dependencies for building the application
RUN apt-get update && apt-get install -y gcc

# Copy the source code into the container
COPY . /app

# Compile the main.go file
RUN GOOS=linux go build -o main ./cmd/app/main.go

# Expose the port the app listens on
EXPOSE 8080

# Command to run the application
CMD ["./main"]