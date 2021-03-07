MIGRATIONS_DIR=infrastructure/datastore/migrations/postgres
POSTGRESQL_URL=postgres://postgres:postgres@localhost:5432/bank?sslmode=disable

dev:
	go mod tidy
	go run cmd/server/main.go

create-migrate:
	migrate create -ext=sql -dir=$(MIGRATIONS_DIR) $(FILE_NAME)

run-migration:
	migrate -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) up

undo-migration:
	migrate -database $(POSTGRESQL_URL) -path $(MIGRATIONS_DIR) down