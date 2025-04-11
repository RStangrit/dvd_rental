FROM golang:1.24 AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -v -o /usr/local/bin/app .

FROM debian:bookworm-slim

COPY --from=builder /usr/local/bin/app /usr/local/bin/app

COPY /assets/fonts/nimbussanl_boldcond.ttf ./assets/fonts/nimbussanl_boldcond.ttf
COPY .env .env
COPY pkg/db/migrations/migration_files ./pkg/db/migrations/migration_files

EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/app"]
