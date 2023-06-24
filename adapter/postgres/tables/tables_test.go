package tables_test

import (
	"context"
	"database/sql"
	"os"
	"testing"

	"github.com/alextanhongpin/core/storage/pg/pgtest"
	"github.com/alextanhongpin/core/test/testutil"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres"
	"github.com/uptrace/bun"
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

type QueryHook struct {
	dump *testutil.SQLDump
}

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	h.dump.Stmt = event.Query
	h.dump.Args = event.QueryArgs
}
