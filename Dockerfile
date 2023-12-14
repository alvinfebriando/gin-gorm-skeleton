FROM golang:1.18.10-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/rest ./cmd/rest/rest.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/migrate ./cmd/migrate/migrate.go
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/seed ./cmd/seed/seed.go

FROM golang:1.18.10-alpine as watcher

RUN apk add --no-cache make
RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

CMD make reload

FROM alpine:3 as migration

WORKDIR /app

RUN apk add --no-cache make

COPY --from=builder /app/bin/migrate /app/bin/
COPY --from=builder /app/bin/seed /app/bin/
COPY --from=builder /app/Makefile /app/

CMD make migration

FROM alpine:3 as dev

WORKDIR /app

COPY --from=builder /app/bin/rest /app/rest

EXPOSE 8080

CMD ./rest
