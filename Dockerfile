# Start from the official Golang image
FROM golang:latest

# Set the Current Working Directory inside the container
WORKDIR /cloud_bee

# Copy the local package files to the container's workspace
COPY . .

# Download all dependencies. They will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Build the server
RUN go build -o server/main ./server

# Build client
RUN go build -o client/main ./client
# RUN go build -o client/main ./client

# Expose port 8080 to the outside world
EXPOSE 8080

# Executable permission for .sh
RUN chmod +x ./run.sh

# Command to run the executable
CMD ["./run.sh"]

# build using this command 
# docker build -t cloudbee . && docker run -p 8080:8080 cloudbee