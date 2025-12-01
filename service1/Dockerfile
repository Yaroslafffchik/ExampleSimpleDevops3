FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod  go.sum ./
RUN apk add --no-cache git
RUN go mod download || true
COPY . .
RUN go build -o /app/service1 main.go

FROM alpine:3.18
WORKDIR /app
RUN apk add --no-cache ca-certificates curl
COPY --from=builder /app/service1 .
RUN chmod +x /app/service1
CMD ["./service1"]