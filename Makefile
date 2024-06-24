gen:
	protoc --go_out=. --go-grpc_out=. sale.proto

export POSTGRES_DB=postgres://postgres:1234@localhost:5432/transactions?sslmode=disable

migfile:
	migrate create -ext sql -dir migrations/ -seq clientstream

up:
	migrate -path migrations -database $(POSTGRES_DB) up

down:
	migrate -path migrations -database $(POSTGRES_DB) down

force:
	migrate -path migrations -database $(POSTGRES_DB) force $(version)

.PHONY: gen migfile up down force
