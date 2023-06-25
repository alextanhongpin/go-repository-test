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
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

func createUser(t *testing.T, db *bun.DB) *tables.User {
	t.Helper()

	ctx := context.Background()
	tbl := tables.NewUser(buntx.New(db))
	u, err := tbl.Create(ctx, "John Appleseed")
	assert.Nil(t, err)
	assert.NotNil(t, u)

	return u
}

func TestProduct(t *testing.T) {
	assert := assert.New(t)

	dump := &testutil.SQLDump{}
	db := pgtest.BunTx(t)
	db.AddQueryHook(&QueryHook{dump: dump})
	user := createUser(t, db)

	ctx := context.Background()
	tbl := tables.NewProduct(buntx.New(db))

	// Create.
	name := "Rainbow Socks"
	desc := "A rainbow colored socks"
	userID := user.ID

	product, err := tbl.Create(ctx, name, desc, userID)
	assert.Nil(err)
	type createProduct struct {
		Args map[string]any
		Rows *tables.Product
	}
	testutil.DumpJSON(t, createProduct{
		Args: map[string]any{
			"Name":   name,
			"Desc":   desc,
			"UserID": userID,
		},
		Rows: product,
	},
		testutil.IgnoreFields("ID", "CreatedAt", "UpdatedAt", "UserID"),
	)

	testutil.DumpPostgres(t,
		dump.WithRows(product),
		testutil.Normalize(),
		testutil.FileName("create_product"),
		testutil.IgnoreFields("$3", "ID", "CreatedAt", "UpdatedAt", "UserID"))

	// Read.
	socks, err := tbl.Find(ctx, product.ID)
	assert.Nil(err)
	assert.Equal(socks, product)

	testutil.DumpPostgres(t,
		dump.WithRows(socks),
		testutil.Normalize(),
		testutil.FileName("find_products"),
		testutil.IgnoreFields("ID", "CreatedAt", "UpdatedAt", "UserID"))

	// Delete.
	err = tbl.Delete(ctx, product.ID)
	assert.Nil(err)

	// Check deleted.
	_, err = tbl.Find(ctx, product.ID)
	assert.NotNil(err)
	assert.True(errors.Is(err, sql.ErrNoRows))
}
