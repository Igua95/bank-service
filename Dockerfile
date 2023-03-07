# Build stage
FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app
# Copy everything inside the /app directory
COPY . . 
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz


# we only need the binary file for the server (we should use multy stage)
# Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY db/migration ./migration
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
RUN chmod +x ./start.sh

EXPOSE 5001
CMD [ "/app/main" ]
# CMD will run as a second argument of entrypoint
ENTRYPOINT [ "/app/start.sh" ]