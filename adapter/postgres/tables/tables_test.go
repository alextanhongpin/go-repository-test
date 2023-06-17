package tables_test

import (
	"database/sql"
	"os"
	"testing"

	_ "embed"

	"github.com/alextanhongpin/core/storage/pg/pgtest"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres"
)

const postgresVersion = "15.1-alpine"

func TestMain(m *testing.M) {
	hook := func(sql *sql.DB) error {
		return postgres.Migrate(pgtest.DSN())
	}
	stop := pgtest.InitDB(pgtest.Tag(postgresVersion), pgtest.Hook(hook))
	code := m.Run()
	stop()
	os.Exit(code)
}
