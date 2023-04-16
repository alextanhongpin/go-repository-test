.PHONY: reset

POSTGRES_PATH := ./adapter/postgres

DBMATE_MIGRATIONS_TABLE := schema_migrations
DBMATE_MIGRATIONS_DIR := $(POSTGRES_PATH)/migrations/
DBMATE_SCHEMA_FILE := $(POSTGRES_PATH)/schemas/schema.sql

DATABASE_URL := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

PGPASSWORD := $(DB_PASS)


db-install:
	@brew install dbmate
	#@brew install libpq # you may need to install this if psql and pg_dump command is not found.


db-help:
	@dbmate --help


sql: ## Creates a new migration
ifndef name
	$(error 'name=<file_name> is required')
else
	@dbmate new $(name)
endif


dump: ## Run the migration
	@dbmate dump


migrate: ## Run the migration
	@dbmate migrate


rollback: ## Undo the last migration
	@dbmate rollback


reset: ## Drop and apply migrations
	-psql -h $(DB_HOST) -p $(DB_PORT) -U $(DB_USER) -d $(DB_NAME) -w -c "select pg_terminate_backend(pid) from pg_stat_activity where datname = 'test'";
	@dbmate drop
	@dbmate up
