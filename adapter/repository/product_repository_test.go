package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-core-microservice/containers"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/go-repository-test/adapter/repository"
	"github.com/stretchr/testify/assert"
)

func TestProductRepository(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	// Setup repository.
	db := containers.PostgresBunDB(t)
	client := buntx.New(db)

	john, socks := testProductRepositorySeed(t, client)

	repo := repository.NewProductRepository(client)

	t.Run("find product succeed", func(t *testing.T) {
		p, err := repo.Find(ctx, socks.ID)
		assert.Nil(err)
		assert.NotNil(p)
		assert.NotNil(p.User, "expect product to have user")
		assert.Equal(repository.NewProduct(socks), p.Product)
		assert.Equal(repository.NewUser(john), p.User)
	})

	t.Run("find product failed", func(t *testing.T) {
		// NOTE: The best way to trigger errors in database is through context
		// cancellation.
		// However, this would only work for Postgres, not MySQL, since the MySQL
		// package doesn't honor context cancellation.
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		p, err := repo.Find(ctx, socks.ID)
		assert.Nil(p)
		assert.True(errors.Is(err, context.Canceled))
	})

	t.Run("list product succeed", func(t *testing.T) {
		p, err := repo.List(ctx)
		assert.Nil(err)
		assert.Equal(1, len(p))
		assert.Equal(repository.NewProduct(socks), p[0])
	})

	t.Run("list product failed", func(t *testing.T) {
		ctx, cancel := context.WithCancel(ctx)
		cancel()

		p, err := repo.List(ctx)
		assert.Nil(p)
		assert.True(errors.Is(err, context.Canceled))
	})
}

func testProductRepositorySeed(t *testing.T, client *buntx.Atomic) (*tables.User, *tables.Product) {
	assert := assert.New(t)
	ctx := context.Background()

	// Create a user.
	userTable := tables.NewUser(client)
	john, err := userTable.Create(ctx, "John Appleseed")
	assert.Nil(err)

	productTable := tables.NewProduct(client)
	socks, err := productTable.Create(ctx, "Socks", "A plain socks", john.ID)
	assert.Nil(err)

	return john, socks
}
