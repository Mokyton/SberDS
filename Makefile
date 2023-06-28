include .env
export

run:
	go run cmd/main.go

build:
	go build cmd/main.go

swag:
	swag init -g cmd/main.go
