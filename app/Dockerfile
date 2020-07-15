# syntax = docker/dockerfile:1-experimental

FROM golangci/golangci-lint:v1.28-alpine AS lint-base

FROM golang:1.14.4-alpine3.12 AS build

WORKDIR /src

ENV CGO_ENABLED=0

COPY --from=lint-base /usr/bin/golangci-lint /usr/bin/

COPY go.* .
RUN go mod download

COPY . .

ARG DB_HOST
ENV DB_HOST ${DB_HOST}

ARG DB_PORT=5432
ENV DB_PORT ${DB_PORT}

ARG DB_USER
ENV DB_USER ${DB_USER}

ARG DB_PASS
ENV DB_PASS ${DB_PASS}

RUN --mount=type=cache,target=/root/.cache/go-build \
    go test -v ./...

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/.cache/golangci-lint \
    golangci-lint run ./...

RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -o /out/app

FROM alpine:3.12
COPY --from=build /out/app /app
ENTRYPOINT ["/app"]
