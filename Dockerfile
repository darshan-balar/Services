# Stage 1: Build the application
FROM golang:1.20 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o service

# Stage 2: Create a lightweight image to run the application
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root/
RUN mkdir config

# Copy the compiled binary from the builder stage
COPY --from=builder /app/service .
COPY config/config.yaml config/config.yaml

# Expose the port on which the application will run
EXPOSE 8081

ARG DEPLOY
ENV DEPLOY=${DEPLOY}

# Command to run the application
CMD ./service ${DEPLOY}
