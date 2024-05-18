FROM golang:1.20 as build

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

FROM scratch

COPY --from=build /main /main

CMD [ "./main" ]
