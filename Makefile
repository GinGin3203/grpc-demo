.DEFAULT_GOAL := build

build:
	docker compose up db_postgres &
	sleep 5 # чтобы база успела запуститься
	docker compose up users_service

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/users.proto

.PHONY: proto build
