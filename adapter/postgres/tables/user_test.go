package tables_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/alextanhongpin/core/storage/pg/pgtest"
	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/testdump/jsondump"
	"github.com/alextanhongpin/testdump/pgdump"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

func init() {
	jsondump.Register(new(tables.User), jsondump.IgnoreFields("ID", "CreatedAt", "UpdatedAt"))
}

func TestCreateUser(t *testing.T) {
	db := pgtest.BunTx(t)
	tbl := newUserTable(t, db)
	u, err := tbl.Create(ctx, "John Appleseed")
	assert.Nil(t, err)
	jsondump.Dump(t, u)
}

func TestFindUser(t *testing.T) {
	db := pgtest.BunTx(t)
	user := createUser(t, db)

	t.Run("success", func(t *testing.T) {
		tbl := newUserTable(t, db)
		john, err := tbl.Find(ctx, user.ID)
		assert.Nil(t, err)
		jsondump.Dump(t, john)
	})

	t.Run("not found", func(t *testing.T) {
		tbl := newUserTable(t, db)
		_, err := tbl.Find(ctx, uuid.New())
		assert.ErrorIs(t, err, sql.ErrNoRows)
	})
}

func TestDeleteUser(t *testing.T) {
	db := pgtest.BunTx(t)
	user := createUser(t, db)

	t.Run("success", func(t *testing.T) {
		tbl := newUserTable(t, db)
		err := tbl.Delete(ctx, user.ID)
		assert.Nil(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		tbl := newUserTable(t, db)
		err := tbl.Delete(ctx, uuid.New())
		assert.Nil(t, err)
	})
}

func createUser(t *testing.T, db *bun.DB) *tables.User {
	t.Helper()

	ctx := context.Background()
	tbl := tables.NewUser(buntx.New(db))
	u, err := tbl.Create(ctx, "John Appleseed")
	assert.Nil(t, err)
	assert.NotNil(t, u)
	store(t.Name(), "users", u.ID, u)

	return u
}

func newUserTable(t *testing.T, db *bun.DB) *tables.UserTable {
	t.Helper()
	db.AddQueryHook(&QueryHook{
		Recorder: pgdump.NewRecorder(t, pgdump.IgnoreArgs("$1")),
	})

	return tables.NewUser(buntx.New(db))
}
