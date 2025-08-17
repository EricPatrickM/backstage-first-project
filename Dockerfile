FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY ./app .

RUN go mod tidy
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -o main ./src

FROM alpine:3.21

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]