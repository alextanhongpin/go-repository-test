package tables

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:t"`

	ID        uuid.UUID `bun:",pk"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserTableMapper interface {
	Find(ctx context.Context, id uuid.UUID) (*User, error)
	Create(ctx context.Context, name string) (*User, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserTable struct {
	client
}

var _ UserTableMapper = (*UserTable)(nil)

func NewUser(client client) *UserTable {
	return &UserTable{
		client: client,
	}
}

func (t *UserTable) Find(ctx context.Context, id uuid.UUID) (*User, error) {
	user := &User{
		ID: id,
	}

	if err := t.DBTx(ctx).
		NewSelect().
		Model(user).
		WherePK().
		Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (t *UserTable) Create(ctx context.Context, name string) (*User, error) {
	user := &User{
		Name: name,
	}

	if _, err := t.DBTx(ctx).
		NewInsert().
		Model(user).
		Column("name").
		Returning("*").
		Exec(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (t *UserTable) Delete(ctx context.Context, id uuid.UUID) error {
	user := &User{
		ID: id,
	}

	_, err := t.DBTx(ctx).
		NewDelete().
		Model(user).
		WherePK().
		Exec(ctx)

	return err
}
