FROM golang:1.19-alpine3.16 AS builder
LABEL build="builder-go-nocode-import-service"
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main ./service/example-kafka

FROM alpine:3.16
RUN apk --no-cache add tzdata curl
WORKDIR /app
COPY --from=builder /app/main /app
CMD ["./main"]