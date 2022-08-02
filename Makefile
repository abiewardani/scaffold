include .env
export


MIGRATION_DIR=tools/migrations
SEEDS_DIR="tools/seeds"

.SILENT: help compose-up compose-down run linter-dotenv mock create_seed

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

compose-up: ### Run docker-compose
	docker-compose up --build -d && docker-compose logs -f

compose-down: ### Down docker-compose
	docker-compose down --remove-orphans

run: ### run
	go mod tidy && go mod download && \
	DISABLE_SWAGGER_HTTP_HANDLER='' GIN_MODE=debug CGO_ENABLED=0 go run -tags migrate ./cmd


linter-dotenv: ### check by dotenv linter
	dotenv-linter

mock: ### run mockery
	mockery --all -r --case snake

## all database related

psql:
	PGPASSWORD=${DATABASE_PASSWORD} psql -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USERNAME} ${DATABASE_NAME}

create_db:
	PGPASSWORD=${DATABASE_PASSWORD} createdb -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USERNAME} ${DATABASE_NAME}
	PGPASSWORD=${DATABASE_PASSWORD} psql -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USERNAME} ${DATABASE_NAME} -c 'CREATE EXTENSION IF NOT EXISTS pg_stat_statements;'
	PGPASSWORD=${DATABASE_PASSWORD} psql -h ${DATABASE_HOST} -p ${DATABASE_PORT} -U ${DATABASE_USERNAME} ${DATABASE_NAME} -c 'CREATE EXTENSION IF NOT EXISTS "uuid-ossp";'

create_migration:
ifndef name
	$(error name is not set)
else
	touch "$(MIGRATION_DIR)/$$(date +"%Y%m%d%H%M%S")_$(name).sql"
endif

migrate_up:
ifndef step
	migrate -source file://$(MIGRATION_DIR) -database "postgresql://${DB_MASTER_USER}:${DB_MASTER_PASSWORD}@${DB_MASTER_HOST}:${DB_MASTER_PORT}/${DB_MASTER_NAME}?sslmode=disable" up 1
else
	migrate -source file://$(MIGRATION_DIR) -database "postgresql://${DB_MASTER_USER}:${DB_MASTER_PASSWORD}@${DB_MASTER_HOST}:${DB_MASTER_PORT}/${DB_MASTER_NAME}?sslmode=disable" up $(step)
endif

migrate_down:
ifndef step
	migrate -source file://$(MIGRATION_DIR) -database "postgresql://${DB_MASTER_USER}:${DB_MASTER_PASSWORD}@${DB_MASTER_HOST}:${DB_MASTER_PORT}/${DB_MASTER_NAME}?sslmode=disable" down 1
else
	migrate -source file://$(MIGRATION_DIR) -database "postgresql://${DB_MASTER_USER}:${DB_MASTER_PASSWORD}@${DB_MASTER_HOST}:${DB_MASTER_PORT}/${DB_MASTER_NAME}?sslmode=disable" down $(step)
endif

migrate_up_all:
	migrate -source file://$(MIGRATION_DIR) -database "postgresql://${DB_MASTER_USER}:${DB_MASTER_PASSWORD}@${DB_MASTER_HOST}:${DB_MASTER_PORT}/${DB_MASTER_NAME}?sslmode=disable" up

migrate_down_all:
	migrate -source file://$(MIGRATION_DIR) -database "postgresql://${DB_MASTER_USER}:${DB_MASTER_PASSWORD}@${DB_MASTER_HOST}:${DB_MASTER_PORT}/${DB_MASTER_NAME}?sslmode=disable" down

migrate_force:
ifndef version
	$(error version is not set)
else
	migrate -source file://$(MIGRATION_DIR) -database "postgresql://${DB_MASTER_USER}:${DB_MASTER_PASSWORD}@${DB_MASTER_HOST}:${DB_MASTER_PORT}/${DB_MASTER_NAME}?sslmode=${DATABASE_SSL}" force $(version)
endif


create_seed:
ifndef name
	$(error name is not set)
else
	touch "$(SEEDS_DIR)/$$(date +"%Y%m%d%H%M%S")_$(name).sql"
endif

run_seeds:
	for file in $(SEEDS_DIR)/*; do PGPASSWORD=${DB_MASTER_PASSWORD} psql -h ${DB_MASTER_HOST} -p ${DB_MASTER_PORT} -U ${DB_MASTER_USER} ${DB_MASTER_NAME} -f "$$file"; done