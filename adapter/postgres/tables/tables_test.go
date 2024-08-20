package tables_test

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"
	"testing"

	"github.com/alextanhongpin/core/storage/pg/pgtest"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres"
	"github.com/alextanhongpin/testdump/pgdump"
	"github.com/uptrace/bun"
)

// Global context.
var ctx = context.Background()

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

// QueryHook logs the sql statement into testdata/
type QueryHook struct {
	Recorder *pgdump.Recorder
}

func (h *QueryHook) BeforeQuery(ctx context.Context, event *bun.QueryEvent) context.Context {
	return ctx
}

func (h *QueryHook) AfterQuery(ctx context.Context, event *bun.QueryEvent) {
	method := fmt.Sprintf("%s_%s", event.Operation(), event.IQuery.GetTableName())
	h.Recorder.Record(method, event.Query, event.QueryArgs...)
}

// mappings store all the nested entities, namespacec by
// the test name.
// This allows access of constructed nested relations .
var mappings sync.Map

// store stores the entity into the namespaced mapping,
// e.g. store(t.Name(), "users", "1", user)
func store(name string, prefix string, key any, val any) {
	m, _ := mappings.LoadOrStore(name, &sync.Map{})
	m.(*sync.Map).Store(fmt.Sprintf("%s:%v", prefix, key), val)
}

// load allows loading the entity from a namespaced mapping
// e.g. load[*tables.User](t.Name(), "users", "1")
func load[T any](name string, prefix string, key any) T {
	var t T
	m, ok := mappings.Load(name)
	if !ok {
		return t
	}
	v, ok := m.(*sync.Map).Load(fmt.Sprintf("%s:%v", prefix, key))
	if !ok {
		return t
	}
	return v.(T)
}
