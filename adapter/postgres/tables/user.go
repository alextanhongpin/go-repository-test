package tables

import (
	"context"
	"time"

	uow "github.com/alextanhongpin/uow/bun"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        uuid.UUID `bun:",pk"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserTable struct {
	conn uow.UOW
}

func NewUser(conn uow.UOW) *UserTable {
	return &UserTable{
		conn: conn,
	}
}

func (u *UserTable) Find(ctx context.Context, id uuid.UUID) (*User, error) {
	user := &User{
		ID: id,
	}

	if err := u.conn.DB(ctx).
		NewSelect().
		Model(user).
		WherePK().
		Scan(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserTable) Create(ctx context.Context, name string) (*User, error) {
	user := &User{
		Name: name,
	}

	if _, err := u.conn.DB(ctx).
		NewInsert().
		Model(user).
		Column("name").
		Returning("*").
		Exec(ctx); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserTable) Delete(ctx context.Context, id uuid.UUID) error {
	user := &User{
		ID: id,
	}

	_, err := u.conn.DB(ctx).
		NewDelete().
		Model(user).
		WherePK().
		Exec(ctx)

	return err
}
