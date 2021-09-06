FROM golang AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -o ./main

FROM alpine

COPY --from=build /app/main /app/main

ENTRYPOINT ["/app/main"]
