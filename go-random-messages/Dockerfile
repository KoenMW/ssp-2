FROM golang:1.25.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o random-messages ./

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/random-messages .

EXPOSE 8080

CMD ["./random-messages"]
