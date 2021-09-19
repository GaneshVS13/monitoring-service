FROM golang:1.15.0

# Set the Current Working Directory inside the container
WORKDIR /app

RUN export GO111MODULE=on

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

#COPY ../monitoring_service_cfg.json ./

# Build the application
RUN go build -o monitoring-service .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./monitoring-service"]