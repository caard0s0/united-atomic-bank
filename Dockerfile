# Build Stage
FROM golang:1.22.5-alpine3.19 AS builder
WORKDIR /app
COPY . . 
RUN go build -o main cmd/main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run Stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY scripts/start.sh .
COPY scripts/wait-for.sh .
COPY internal/database/migrations ./migrations

EXPOSE 80
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]