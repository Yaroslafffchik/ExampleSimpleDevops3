FROM golang:1.25-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download
COPY . .
RUN go build -o /app/service2 main.go

FROM alpine:3.21
WORKDIR /app
RUN apk add --no-cache curl
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/service2 .
RUN apk add --no-cache curl
RUN chmod +x /app/service2
CMD ["./service2"]