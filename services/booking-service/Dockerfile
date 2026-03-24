FROM golang:1.26 AS build

WORKDIR /build
COPY . /build

RUN go test -mod=vendor ./...
RUN GOOS=linux  GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -o booking-service

FROM alpine:3 AS delpoy

RUN apk update && apk upgrade \
    && apk add ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /app

COPY --from=build /build/booking-service /app/

ENTRYPOINT ["./booking-service"]