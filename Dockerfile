FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .
RUN go build -o main main.go

FROM alpine:3.19

COPY --from=builder /app /app

WORKDIR /app

CMD ["./main"]