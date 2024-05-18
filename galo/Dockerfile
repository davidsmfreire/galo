FROM alpine

RUN apk add --no-cache build-base

WORKDIR /app

COPY main.c .

RUN gcc -o /main main.c
