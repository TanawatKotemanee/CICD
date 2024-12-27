# Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go

# Runner
FROM alpine:latest

WORKDIR /app/

COPY --from=builder /app/main .

# Override the port when running the container.
ENV PORT=8080

EXPOSE $PORT

CMD ["./main"]