FROM swift as build

COPY main.swift .

RUN swiftc main.swift -o /main

FROM swift:slim

COPY --from=build /main /main

CMD [ "./main" ]
