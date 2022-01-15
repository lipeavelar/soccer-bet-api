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

.PHONY: create-migration
create-migration: ## create migration with goose (need to install goose first), usage: make create-migration SBM_NAME="migratrion_name"
	@echo "Creating migration"
ifeq ($(origin SBM_NAME),undefined)
	@echo "Migration name is required, check usage with 'make help'"
	exit 1
endif
	goose -dir=./db/migrations/ create $(SBM_NAME) sql

.PHONY: run
run: ## Run API server and database
	docker-compose up -d