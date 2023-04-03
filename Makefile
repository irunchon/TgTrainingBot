include .env
export

.PHONY: all

all:
	go run cmd/bot/main.go