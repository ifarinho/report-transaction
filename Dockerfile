# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o goapp ./cmd/app/

FROM alpine as runtime

COPY --from=build /app /goapp

ENTRYPOINT ["/goapp/goapp"]
