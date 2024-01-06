package tables_test

import (
	"database/sql"
	"log"
	"strings"
	"testing"
	"time"

	"github.com/alextanhongpin/core/storage/pg/pgtest"
	"github.com/alextanhongpin/core/test/testutil"
	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

func TestCreateProduct(t *testing.T) {
	db := pgtest.BunTx(t)
	db.AddQueryHook(&QueryHook{
		t: t,
		opts: []testutil.SQLOption{
			testutil.IgnoreArgs("$3"),
		},
	})

	t.Run("success", func(t *testing.T) {
		p := createProduct(t, db)
		testutil.DumpJSON(t, p, testutil.IgnoreFields("ID", "CreatedAt", "UpdatedAt", "UserID"))
	})
}

func TestFindProduct(t *testing.T) {
	db := pgtest.BunTx(t)
	p := createProduct(t, db)

	t.Run("success", func(t *testing.T) {
		tbl := newProductTable(t, db)
		res, err := tbl.Find(ctx, p.ID)
		assert.Nil(t, err)
		testutil.DumpJSON(t, res, testutil.IgnoreFields("ID", "CreatedAt", "UpdatedAt", "UserID"))
	})

	t.Run("not found", func(t *testing.T) {
		tbl := newProductTable(t, db)
		_, err := tbl.Find(ctx, -1)
		assert.ErrorIs(t, err, sql.ErrNoRows)
	})
}

func TestDeleteProduct(t *testing.T) {
	db := pgtest.BunTx(t)
	p := createProduct(t, db)

	t.Run("success", func(t *testing.T) {
		tbl := newProductTable(t, db)
		err := tbl.Delete(ctx, p.ID)
		assert.Nil(t, err)
	})

	t.Run("not found", func(t *testing.T) {
		tbl := newProductTable(t, db)
		err := tbl.Delete(ctx, -1)
		assert.Nil(t, err)
	})
}

func createProduct(t *testing.T, db *bun.DB, options ...string) *tables.Product {
	t.Helper()

	p := newProduct(options...)
	if p.UserID == uuid.Nil {
		u := createUser(t, db)
		p.UserID = u.ID
	}

	tbl := tables.NewProduct(buntx.New(db))
	res, err := tbl.Create(ctx, p.Name, p.Description, p.UserID)
	if err != nil {
		t.Fatalf("failed to create product: %v", err)
	}
	store(t.Name(), "products", res.ID, res)

	return res
}

func newProduct(options ...string) *tables.Product {
	p := &tables.Product{
		ID:          0,
		Name:        "table",
		Description: "a wooden table",
		UserID:      uuid.Nil,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	for _, o := range options {
		if strings.HasPrefix(o, "with_user_id:") {
			_, userID, _ := strings.Cut(o, "with_user_id:")
			p.UserID = uuid.MustParse(userID)
			continue
		}

		switch o {
		default:
			log.Fatalf("unknown product option: %s", o)
		}
	}
	return p
}

func newProductTable(t *testing.T, db *bun.DB) *tables.ProductTable {
	t.Helper()
	db.AddQueryHook(&QueryHook{
		t: t,
		opts: []testutil.SQLOption{
			testutil.IgnoreArgs("$3"),
		},
	})

	return tables.NewProduct(buntx.New(db))
}
