# Stage 1: Build the Go application
FROM golang:1.21 AS builder

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Stage 2: Create a minimal scratch container
FROM scratch

# Copy the binary from the builder stage into the scratch container
COPY --from=builder /app/app /app

# Expose port 3000 which the application listens on
EXPOSE 3000

# Set environment variables if necessary (e.g., FILE_PATH and GREETING)
# ENV FILE_PATH=/custom/path
# ENV GREETING=HelloCustom

# Run the Go binary
ENTRYPOINT ["/app"]
