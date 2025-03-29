FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod . 
COPY go.sum . 
RUN go mod download

COPY . . 

RUN go build -o main ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

CMD ["/app/main"]