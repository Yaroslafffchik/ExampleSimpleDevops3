FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download || true
COPY . .
RUN go build -o /app/service2 main.go

FROM alpine:3.18
WORKDIR /app
RUN apk add --no-cache ca-certificates curl
COPY --from=builder /app/service2 .
RUN chmod +x /app/service2
CMD ["./service2"]