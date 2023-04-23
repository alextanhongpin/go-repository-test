package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-core-microservice/containers"
	"github.com/alextanhongpin/go-repository-test/adapter/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepository(t *testing.T) {
	assert := assert.New(t)

	// Setup repository.
	db := containers.PostgresBunDB(t)
	repo := repository.NewAuthRepository(buntx.New(db))

	name := "John Appleseed"
	ctx := context.Background()

	// Create a user.
	user, err := repo.CreateUser(ctx, name)
	assert.Nil(err)

	assert.NotNil(user)
	assert.NotEqual(uuid.Nil, user.ID)
	assert.Equal(name, user.Name)
	assert.Equal(false, user.CreatedAt.IsZero())
	assert.Equal(false, user.UpdatedAt.IsZero())

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
