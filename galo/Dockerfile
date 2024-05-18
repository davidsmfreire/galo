FROM gcc:latest as build

COPY main.c .

RUN gcc -static -o main main.c

FROM scratch

COPY --from=build /main /main

CMD [ "./main" ]
