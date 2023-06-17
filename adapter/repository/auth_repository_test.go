package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/alextanhongpin/core/storage/pg/pgtest"
	"github.com/alextanhongpin/core/test/testutil"
	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-repository-test/adapter/repository"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepository(t *testing.T) {
	assert := assert.New(t)

	// Setup repository.
	db := pgtest.BunTx(t)
	repo := repository.NewAuthRepository(buntx.New(db))

	name := "John Appleseed"
	ctx := context.Background()

	// Create a user.
	user, err := repo.CreateUser(ctx, name)
	assert.Nil(err)

	// Dump.
	testutil.DumpJSON(t, user, testutil.IgnoreFields("ID", "CreatedAt", "UpdatedAt"))

	// Find the created user.
	john, err := repo.FindUser(ctx, user.ID)
	assert.Equal(user, john)
	assert.Nil(err)

	// Delete the user.
	err = repo.Delete(ctx, user.ID)
	assert.Nil(err)

	// Verify the user has been deleted.
	_, err = repo.FindUser(ctx, user.ID)
	assert.True(errors.Is(err, sql.ErrNoRows), "expect user to be not found")
}
