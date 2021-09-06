FROM golang AS build

WORKDIR /app

COPY ./main.go ./main.go

RUN CGO_ENABLED=0 go build -o ./main ./main.go

FROM alpine

COPY --from=build /app/main /app/main

ENTRYPOINT ["/app/main"]
