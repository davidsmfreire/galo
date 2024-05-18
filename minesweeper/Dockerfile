FROM golang:1.20

WORKDIR /app

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go
