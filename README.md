# go-repository-test


- always handle between db and tx in the usecase layer
- the `tables` (or `storage`, or any other relevant name to your liking) represents the mapping between the schema in the db to the golang code
- each usecase will have __one__ repository
- this makes mocking easier, instead of having multiple repositories that is actually calling one table, this structure places the table mapping in the `tables` layer, and the `repository` layer acts as a facade that calls different `tables`
- note that the repository layer can also
	- call other third-party client apis
	- do caching transparently (e.g. check data exists in cache, otherwise query db results and cache them before returning)
	- maps table-specific or external data types to __domain__ types (this is very important)
- migrations can be versioned, or done declaratively
	- using atlas as the tool, this does not have the concept of rollback migration
	- edit tables in `schemas`, and run `make atlas-diff-apply` to update the db state
	- for some usecases however, atlas cannot diff the changes
	- so for extensions, triggers and other dlls, we run manual migration
	- always run manual migration before declarative migration


## Running migration

With Atlas (note: this is not ideal, use dbmate instead):
```bash
# Run versioned migrations (triggers, extensions, seeding data).
$ make atlas-migrate

# Run declarative migrations (schema changes).
$ make atlas-diff-apply

# Note that due to the sequence on migrations, the diff may have dependencies on extensions etc, so we need to re-run this again. Some triggers can only be created once the tables are created too.
# One alternative is to put the basic table creation in the versioned migration too and adding the trigger there.
$ make atlas-migrate
```

With dbmate:

```bash
$ make sql name=create_table...
$ make migrate
$ make rollback
$ make reset
```


## Other migration options

declarative:
- [migra](https://github.com/djrobstep/migra) for postgres, example [here](https://github.com/alextanhongpin/ask-and-answer) and [here](https://github.com/alextanhongpin/go-food)
- [skeema](https://www.skeema.io/) for mysql
- [dbmate](https://github.com/alextanhongpin/go-passport)
