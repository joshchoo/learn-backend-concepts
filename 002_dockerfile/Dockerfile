FROM golang AS build

WORKDIR /app

COPY ./main.go ./main.go

RUN go build -o ./main ./main.go

FROM alpine

COPY --from=build /app/main /app/main

CMD ["/app/main"]