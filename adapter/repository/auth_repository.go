package repository

import (
	"context"
	"fmt"

	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/go-repository-test/domain"
	"github.com/google/uuid"
)

type AuthRepository struct {
	users tables.UserTableMapper
}

func NewAuthRepository(db atomic) *AuthRepository {
	return &AuthRepository{
		users: tables.NewUser(db),
	}
}

func (r *AuthRepository) FindUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	user, err := r.users.Find(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("users.Find: %w", err)
	}

	return NewUser(user), nil
}

func (r *AuthRepository) CreateUser(ctx context.Context, name string) (*domain.User, error) {
	user, err := r.users.Create(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("users.Create: %w", err)
	}

	return NewUser(user), nil
}

func (r *AuthRepository) Delete(ctx context.Context, id uuid.UUID) error {
	err := r.users.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("users.Delete: %w", err)
	}

	return nil
}
