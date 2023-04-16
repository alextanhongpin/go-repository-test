package tables_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/alextanhongpin/go-core-microservice/containers"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/uow/bun"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	db := containers.PostgresBunDB(t)
	uow := bun.New(db)
	userTable := tables.NewUser(uow)

	// Create.
	user, err := userTable.Create(ctx, "john appleseed")
	assert.Nil(err)
	assert.NotNil(user)
	assert.True(user.ID != uuid.Nil)

	// Read.
	john, err := userTable.Find(ctx, user.ID)
	assert.Nil(err)
	assert.Equal(john, user)

	// Delete.
	err = userTable.Delete(ctx, user.ID)
	assert.Nil(err)

	// Check deleted.
	_, err = userTable.Find(ctx, user.ID)
	assert.NotNil(err)
	assert.True(errors.Is(err, sql.ErrNoRows))
}
