FROM golang:1.20 AS builder
WORKDIR /app/
COPY . .
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./app ./cmd/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
ENTRYPOINT [ "/app/app" ]