package repository

import (
	"github.com/alextanhongpin/go-repository-test/adapter/postgres/tables"
	"github.com/alextanhongpin/go-repository-test/domain"
)

func NewProduct(p *tables.Product) *domain.Product {
	return &domain.Product{
		ID:          p.ID,
		Name:        p.Name,
		Description: p.Description,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func NewProductOwner(p *tables.Product) *domain.ProductOwner {
	return &domain.ProductOwner{
		Product: NewProduct(p),
		User:    NewUser(p.User),
	}
}
