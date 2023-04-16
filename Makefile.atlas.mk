pg_conn := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable
pg_conn_dev := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/dev?sslmode=disable


TESTDATA_PATH := $(PWD)/adapter/postgres/tables/testdata

# The path that contains all the schema files.
SCHEMA_PATH := $(PWD)/adapter/postgres/schemas

# The path that contains all the versioned migration files.
MIGRATION_PATH := $(PWD)/adapter/postgres/migrations

ATLAS_IMAGE := arigaio/atlas:0.10.0-alpine
ATLAS := docker run --rm --net=host --volume=$(SCHEMA_PATH):/schemas --volume=$(MIGRATION_PATH):/migrations $(ATLAS_IMAGE)


# Alternative is to install this, but using the docker image is preferred.
atlas-install:
	@go install ariga.io/atlas/cmd/atlas@latest


atlas-help:
	@$(ATLAS) help


atlas-hash:
	@$(ATLAS) migrate hash --dir=file://schemas


# Computes the diff between the target db with the local schemas.
# We exclude the table generated by ATLAS to store the migration revision, atlas_schema_revisions
atlas-diff: atlas-hash
	@$(ATLAS) schema diff \
		--from $(pg_conn) \
		--to file://schemas \
		--dev-url $(pg_conn_dev) \
		--exclude 'atlas_schema_revisions' > diff.sql
	@cat diff.sql
	@rm diff.sql


# Applies the migration.
atlas-diff-apply: atlas-hash
	@$(ATLAS) schema apply \
		--url $(pg_conn) \
		--to file://schemas \
		--dev-url $(pg_conn_dev) \
		--exclude 'atlas_schema_revisions' \
		--auto-approve


atlas-diff-apply-dry-run: atlas-hash
	@$(ATLAS) schema apply \
		--url $(pg_conn) \
		--to file://schemas \
		--dev-url $(pg_conn_dev) \
		--exclude 'atlas_schema_revisions' \
		--dry-run


# Check database schemas.
# NOTE: It is hard to migrate during test, so we always generate the schema.
atlas-inspect:
	@mkdir -p $(TESTDATA_PATH)
	@$(ATLAS) schema inspect --url $(pg_conn) --schema public --format '{{ sql . }}' > $(TESTDATA_PATH)/baseline.sql

atlas-new: atlas-hash
ifndef name
	$(error 'name is required')
else
	@$(ATLAS) migrate new $(name)\
		--dir file://migrations \
		--dir-format atlas
endif


atlas-hash-migrate:
	@$(ATLAS) migrate hash --dir=file://migrations


atlas-migrate: atlas-hash-migrate
	@$(ATLAS) migrate apply $(name)\
		--dir file://migrations \
		--url $(pg_conn) $(n)
