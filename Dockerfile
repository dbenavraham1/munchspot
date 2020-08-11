# Build stage
FROM golang AS builder

# Working directory
WORKDIR /app

# Get dependencies
COPY go.* ./
COPY app.yml ./
RUN go mod download

# Copy project to working directory
COPY . ./

# Build staged container
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server

# Final Stage
FROM alpine:latest

# Install ca-certificates bundle inside the docker image
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/app.yml /app.yml
RUN sed -i 's/host: localhost/host:/g' /app.yml
COPY --from=builder /app/server /server

# Run the web service on container startup.
CMD ["./server"]
