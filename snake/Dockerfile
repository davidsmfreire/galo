FROM rust:slim as build

COPY main.rs .

RUN rustc main.rs -o /main

FROM gcr.io/distroless/cc-debian12

COPY --from=build /main /main

CMD [ "./main" ]
