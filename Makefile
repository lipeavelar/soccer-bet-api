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
	goose -dir ./database/migrations/ postgres 'host=localhost port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable' up

.PHONY: reset-migration
reset-migration: ## Reset last migration
	@echo "Migrating from old Makefile to new Makefile"
	goose -dir ./database/migrations/ postgres 'host=localhost port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable' down

.PHONY: migration-status
migration-status: ## Run migrations with goose (need to install goose first)
	@echo "Migrating from old Makefile to new Makefile"
	goose -dir ./database/migrations/ postgres 'host=localhost port=$(DB_PORT) user=$(DB_USER) password=$(DB_PASS) dbname=$(DB_NAME) sslmode=disable' status

.PHONY: create-migration
create-migration: ## Create migration with goose (need to install goose first), usage: make create-migration SBM_NAME="migratrion_name"
	@echo "Creating migration"
ifeq ($(origin SBM_NAME),undefined)
	@echo "Migration name is required, check usage with 'make help'"
	exit 1
endif
	goose -dir=./database/migrations/ create $(SBM_NAME) sql

.PHONY: start
start: ## Run API server and database
	docker-compose up

.PHONY: stop
stop: ## Stops API server and database
	docker-compose down

.PHONY: check-vuln
check-vuln: ## Check vulnerabilities on code base
	govulncheck ./...
