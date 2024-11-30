FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app
COPY . .

RUN mkdir -p bin/
RUN go build -o bin/gateway cmd/gateway/main.go
RUN go build -o bin/server cmd/server/main.go

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/bin/gateway /app/bin/gateway
COPY --from=builder /app/bin/server /app/bin/server
