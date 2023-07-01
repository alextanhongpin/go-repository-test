package tables_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/alextanhongpin/core/storage/pg/pgtest"
	"github.com/alextanhongpin/core/test/testutil"
	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	// Create a new dump
	dump := &testutil.SQLDump{}
	db := pgtest.BunTx(t)

	// Pass it to query hook.
	db.AddQueryHook(&QueryHook{dump: dump})
	tbl := tables.NewUser(buntx.New(db))

	// Create.
	name := "john appleseed"
	user, err := tbl.Create(ctx, name)
	assert.Nil(err)
	type createUser struct {
		Args map[string]any
		Rows *tables.User
	}

	testutil.DumpJSON(t, createUser{
		Args: map[string]any{
			"Name": name,
		},
		Rows: user,
	},
		testutil.IgnoreFields("ID", "CreatedAt", "UpdatedAt"),
	)
	assert.NotNil(user)
	assert.True(user.ID != uuid.Nil)

	testutil.DumpPostgres(t,
		dump.WithResult(user),
		testutil.IgnoreRows("ID", "CreatedAt", "UpdatedAt"),
	)

	// Read.
	john, err := tbl.Find(ctx, user.ID)
	assert.Nil(err)
	assert.Equal(john, user)

	// Delete.
	err = tbl.Delete(ctx, user.ID)
	assert.Nil(err)

	// Check deleted.
	_, err = tbl.Find(ctx, user.ID)
	assert.NotNil(err)
	assert.True(errors.Is(err, sql.ErrNoRows))
}
