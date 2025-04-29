MIGRATOR_CMD=go run currency/cmd/migrator/main.go
DB_URL=postgres://postgres:postgres@localhost:5432/currency?sslmode=disable

.PHONY: migrate-up migrate-down migrate-version build-migration


refresh-migration:migrate-down migrate-up

migrate-up:
	$(MIGRATOR_CMD) -action=up

migrate-down:
	$(MIGRATOR_CMD) -action=down

migrate-version:
	@read -p "Enter version to migrate to: " v && $(MIGRATOR_CMD) -action=version -version=$$v
