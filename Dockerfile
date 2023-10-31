# syntax=docker/dockerfile:1

FROM golang:1.21-alpine as build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . ./

RUN go build -o goapp ./cmd/app/

FROM alpine as runtime

COPY --from=build /app /goapp

COPY scripts/docker_entrypoint.sh .

ENTRYPOINT ["sh", "/docker_entrypoint.sh", "-filename", "-account"]
