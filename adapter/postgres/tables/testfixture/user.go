package testfixture

import (
	"context"
	"testing"
	"time"

	"github.com/alextanhongpin/dbtx/buntx"
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/google/uuid"
)

// client represents either a db or tx
// client.
type client interface {
	DBTx(ctx context.Context) buntx.DBTX
}

func CreateUser(t *testing.T, ctx context.Context, c client, variants ...string) *tables.User {
	t.Helper()

	u := NewUser(variants...)

	user, err := tables.NewUser(c).Create(ctx, u.Name)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	store(t.Name(), "users", user.ID, user)

	return user
}

func NewUser(variants ...string) *tables.User {
	return &tables.User{
		ID:        uuid.Nil,
		Name:      "john",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
