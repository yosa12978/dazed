FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o bin/dazed ./cmd/dazed/main.go

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/bin .
COPY --from=builder /app/config.yaml .
COPY --from=builder /app/migrations ./migrations

RUN apk --update --no-cache add curl

EXPOSE 80
ENTRYPOINT ["./dazed"]