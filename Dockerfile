# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:alpine

# Add Maintainer Info
LABEL maintainer="Domingo Sanz Marti <domingosanzmarti@gmail.com>"

RUN GOCACHE=OFF

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
