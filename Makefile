include .env
export

### Help example https://stackoverflow.com/a/64996042/1606920
.PHONY: help
help:
	@echo 'Usage: make [options] [target]'
	@egrep -h '\s##\s' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-30s\033[0m %s\n", $$1, $$2}'

.PHONY: migrate
migrate: ## Run migrations with goose (need to install goose first)
	@echo "Migrating from old Makefile to new Makefile"
	goose -dir ./db/migrations/ postgres 'host=localhost port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable' up

.PHONY: migrate-status
migrate-status: ## Run migrations with goose (need to install goose first)
	@echo "Migrating from old Makefile to new Makefile"
	goose -dir ./db/migrations/ postgres 'host=localhost port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable' status

.PHONY: run-api
run-api: ## Run API server and database
	docker-compose up -d