package tables_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-core-microservice/containers"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/stretchr/testify/assert"
	"github.com/uptrace/bun"
)

func createUser(t *testing.T, db *bun.DB) *tables.User {
	userTable := tables.NewUser(buntx.New(db))
	u, err := userTable.Create(context.Background(), "John Appleseed")
	assert.Nil(t, err)
	assert.NotNil(t, u)
	return u
}

func TestProduct(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	db := containers.PostgresBunDB(t)
	user := createUser(t, db)

	productTable := tables.NewProduct(buntx.New(db))

	// Create.
	name := "Rainbow Socks"
	desc := "A rainbow colored socks"
	userID := user.ID

	product, err := productTable.Create(ctx, name, desc, userID)
	assert.Nil(err)
	assert.NotNil(product)
	assert.True(product.ID != 0)
	assert.Equal(name, product.Name)
	assert.Equal(desc, product.Description)
	assert.Equal(userID, product.UserID)

	// Read.
	socks, err := productTable.Find(ctx, product.ID)
	assert.Nil(err)
	assert.Equal(socks, product)

	// Delete.
	err = productTable.Delete(ctx, product.ID)
	assert.Nil(err)

	// Check deleted.
	_, err = productTable.Find(ctx, product.ID)
	assert.NotNil(err)
	assert.True(errors.Is(err, sql.ErrNoRows))
}
