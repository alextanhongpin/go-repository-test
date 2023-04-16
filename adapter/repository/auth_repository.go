package repository

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/go-repository-test/domain"
	uow "github.com/alextanhongpin/uow/bun"
	"github.com/google/uuid"
)

type AuthRepository struct {
	conn      uow.UOW
	userTable *tables.UserTable
}

func NewAuth(conn uow.UOW) *AuthRepository {
	return &AuthRepository{
		conn:      conn,
		userTable: tables.NewUser(conn),
	}
}

func (r *AuthRepository) FindUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.userTable.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("userTable.Find: %w", err)
	}

	return newUser(user), nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, name string) (*domain.User, error) {
	user, err := r.userTable.Create(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("userTable.Create: %w", err)
	}

	return newUser(user), nil
}

func (r *AuthRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.userTable.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("userTable.Delete: %w", err)
	}

	return nil
}
