# syntax=docker/dockerfile:1

## Build
FROM golang:1.19-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /jarvis

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /jarvis /jarvis

USER nonroot:nonroot

EXPOSE 80

ENTRYPOINT ["/jarvis"]