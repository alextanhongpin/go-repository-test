// file auth.go contains the logic to performing mapping
// from tables types to domain types.
// As there could be multiple conversion to the same domain type, we can
// standardize the naming as such:
// - If the conversion is from the same type, e.g. tables.User to domain.User,
// we name the method newUser
// - If the conversion is from different type, e.g. tables.User to
// domain.Credentials, we name the method newCredentialsFromUser
package repository

import (
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/go-repository-test/domain"
)

func NewUser(u *tables.User) *domain.User {
	return &domain.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
