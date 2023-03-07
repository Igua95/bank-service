# Build stage
FROM golang:1.19-alpine3.17 AS builder
WORKDIR /app
# Copy everything inside the /app directory
COPY . . 
RUN go build -o main main.go

# we only need the binary file for the server (we should use multy stage)
# Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 5001
CMD [ "/app/main" ]