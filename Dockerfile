# Builder
FROM golang:1.20-bullseye AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o /bspliter

# Prod
FROM debian:bullseye-slim AS prod
COPY --from=builder /gotodo /gotodo
USER nobody:nogroup
ENTRYPOINT ["/bspliter"]

# Dev
FROM golang:1.20-bullseye AS dev
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
ENTRYPOINT ["air"]
