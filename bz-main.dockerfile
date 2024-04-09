# Stage 1: build the Go binary
FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Stage 2: Copy the Go binary to a minimal base image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

# Add the wait-for script
COPY wait-for-redis.sh /root/
RUN chmod +x /root/wait-for-redis.sh

# Modify the CMD to run the wait-for script before the main application
CMD ["/bin/sh", "-c", "/root/wait-for-redis.sh && ./main"]