FROM golang AS build

WORKDIR /app

COPY ./main.go ./main.go
COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

RUN CGO_ENABLED=0 go build -o ./main ./main.go

FROM alpine

COPY --from=build /app/main /app/main

ENTRYPOINT ["/app/main"]
