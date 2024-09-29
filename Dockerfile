FROM golang:1.23

WORKDIR /app
COPY . .

RUN mkdir -p bin/
RUN go build -o bin/gateway cmd/gateway/main.go
RUN go build -o bin/server cmd/server/main.go
