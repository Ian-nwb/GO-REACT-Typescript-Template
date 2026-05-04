# Stage 1: Build the binary
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the app - pointing to our entry point
RUN go build -o main ./cmd/api/main.go

# Stage 2: Run the binary
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]