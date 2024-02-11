# Use the official Golang image as base
FROM golang:1.22 as builder

# Set the working directory
WORKDIR /backend

# Set environment variables for cross-compilation
ENV GOOS=linux
ENV GOARCH=amd64

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application for Linux
RUN CGO_ENABLED=0 go build -o /backend ./cmd/api/main.go

# Use a minimal Alpine image as the final base
FROM alpine:latest

# Install CA certificates required for HTTPS
RUN apk --no-cache add ca-certificates

# Copy the built binary from the previous stage
COPY --from=builder /backend /

# Expose the port on which the application will run
EXPOSE 8080

# Set the entrypoint for the container
ENTRYPOINT ["/backend"]

